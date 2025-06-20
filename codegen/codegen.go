package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/gertd/go-pluralize"
	"github.com/rs/zerolog/log"
)

// 生产CRUD代码
func Gen(model string, ins interface{}) {
	param := "c *gin.Context"

	funcAdd := fmt.Sprintf("%s%s", model, "AddHandler")
	funcList := fmt.Sprintf("%s%s", model, "ListHandler")
	funcDelete := fmt.Sprintf("%s%s", model, "DeleteHandler")
	funcUpdate := fmt.Sprintf("%s%s", model, "UpdateHandler")

	pluralize := pluralize.NewClient()
	lowerModel := strings.ToLower(model)
	pkgName := pluralize.Plural(lowerModel)

	methods := []string{}
	insT := reflect.TypeOf(ins)
	for i := 0; i < insT.NumMethod(); i++ {

		method := insT.Method(i).Name
		if strings.HasPrefix(method, "Set") && !strings.Contains(method, "Nillable") && method != "SetCreatedAt" && method != "SetUpdatedAt" && method != "SetDeleted" {
			methods = append(methods, method)
		}
	}

	assignments := List(Id(lowerModel), Err()).Op(":=").Id("db.Client." + model + ".Create().").Line()
	assignmentsForUpdate := List(Id(lowerModel), Err()).Op(":=").Id("db.Client." + model + ".UpdateOneID(form.ID).").Line()
	for _, v := range methods {
		assignments = assignments.Id(v + "(form." + strings.TrimPrefix(v, "Set") + ").").Line()
		assignmentsForUpdate = assignmentsForUpdate.Id(v + "(form." + strings.TrimPrefix(v, "Set") + ").").Line()
	}
	assignments = assignments.Id(`SetTenantId(form.TenantId).
	SetTenantName(form.TenantName).`)
	assignments = assignments.Id("SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).").Line().
		Id("Save(c.Request.Context())")
	assignmentsForUpdate = assignmentsForUpdate.Id("SetUpdatedAt(time.Now().Unix()).").Line().
		Id("Save(c.Request.Context())")

	f := NewFile(pkgName)

	//  Id("import (").Line().
	// 	Id("cattleai/db").Line().
	// 	Id("cattleai/ent").Line().
	// 	Id("cattleai/ent/tenant").Line().
	// 	Id("cattleai/pkg/paging").Line().
	// 	Id("cattleai/pkg/params").Line().
	// 	Id("cattleai/resp").Line().
	// 	Id("fmt").Line().
	// 	Id("net/http").Line().
	// 	Id("time").Line().
	// 	Id("github.com/gin-gonic/gin").Line().
	// 	Id("github.com/rs/zerolog/log").Line().
	// 	Id(")")

	f.Func().Id(funcAdd).Params(Id(param)).Block(
		Id("form").Op(":=").Id("&ent."+model+"{}"),
		If(Err().Op(":=").Id("c.Bind(form)"), Err().Op("!=").Id("nil")).Block(
			Id("log.Error().Msg(err.Error())"),
			Id("return"),
		),
		Id("log.Debug().Msg(form.String())"),
		assignments,
		If(Err().Op("!=").Id("nil")).Block(
			Id("log.Error().Msg(err.Error())"),
			Id("c.Status(http.StatusInternalServerError)"),
			Id("return"),
		),
		Id("c.JSON(http.StatusOK, resp.Success("+lowerModel+"))"),
	).Line()

	f.Func().Id(funcList).Params(Id(param)).Block(
		Id("listParams").Op(":=").Id("&params.ListParams{}"),
		If(Err().Op(":=").Id("c.BindQuery(listParams)"), Err().Op("!=").Id("nil")).Block(
			Id("log.Error().Msg(err.Error())"),
			Id("return"),
		),
		Id("page").Op(":=").Id("listParams.Paging"),
		Id(`listParams.Level = c.MustGet("level").(int)`),
		Id("where").Op(":=").Id("Where(listParams)"),
		List(Id("totalCount"), Err()).Op(":=").Id("db.Client."+model+".Query().Where(where).Count(c.Request.Context())"),
		If(Err().Op("!=").Id("nil")).Block(
			Id("log.Error().Msg(err.Error())"),
			Id("c.Status(http.StatusInternalServerError)"),
			Id("return"),
		),
		// TODO
		List(Id(pkgName), Err()).Op(":=").Id("db.Client."+model+".Query().Where(where).Order(ent.Desc("+lowerModel+".FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())"),
		If(Err().Op("!=").Id("nil")).Block(
			Id("log.Error().Msg(err.Error())"),
			Id("c.Status(http.StatusInternalServerError)"),
			Id("return"),
		),
		Id("page.TotalCount").Op("=").Id("totalCount"),
		Id("pageData").Op(":=").Id("paging.PageData{").Line().Id("Data: "+pkgName+",").Line().
			Id("Paging: page,").Line().Id("}"),
		Id("c.JSON(http.StatusOK, resp.Success(pageData))"),
	).Line()

	f.Func().Id(funcDelete).Params(Id(param)).Block(
		Id("id").Op(":=").Id("&params.Id{}"),
		If(Err().Op(":=").Id("c.BindUri(id)"), Err().Op("!=").Id("nil")).Block(
			Id("log.Error().Msg(err.Error())"),
			Id("return"),
		),
		Id(`log.Debug().Msg(fmt.Sprintf("%+v", id))`),
		List(Id(lowerModel), Err()).Op(":=").Id("db.Client."+model+".DeleteOneID(int(id.Id)).Exec(c.Request.Context())"),
		If(Err().Op("!=").Id("nil")).Block(
			Id("log.Error().Msg(err.Error())"),
			Id("c.Status(http.StatusInternalServerError)"),
			Id("return"),
		),
		Id("c.JSON(http.StatusOK, resp.Success("+lowerModel+"))"),
	).Line()

	f.Func().Id(funcUpdate).Params(Id(param)).Block(
		Id("form").Op(":=").Id("&ent."+model+"{}"),
		If(Err().Op(":=").Id("c.Bind(form)"), Err().Op("!=").Id("nil")).Block(
			Id("log.Error().Msg(err.Error())"),
			Id("return"),
		),
		Id("log.Debug().Msg(form.String())"),
		assignmentsForUpdate,
		If(Err().Op("!=").Id("nil")).Block(
			Id("log.Error().Msg(err.Error())"),
			Id("c.Status(http.StatusInternalServerError)"),
			Id("return"),
		),
		Id("c.JSON(http.StatusOK, resp.Success("+lowerModel+"))"),
	).Line()

	if err := os.Mkdir("../pkg/"+pkgName, 0755); err != nil {
		log.Warn().Msg(err.Error())
	}

	err := f.Save("../pkg/" + pkgName + "/g_handlers.go")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", f)

	whereF := NewFile(pkgName)
	whereF.Func().Id("Where").Params(Id("listParams *params.ListParams")).Id("predicate."+model).Block(
		Id("wheres").Op(":=").Id("[]predicate."+model+"{"+lowerModel+".Deleted(0)}"),
		If(Id("listParams.Q")).Op("!=").Id(`""`).Block(
			Id("wheres").Op("=").Id("append(wheres, "+lowerModel+".NameContains(listParams.Q))"),
		),
		Id("wheres").Op("=").Id("append(wheres, "+lowerModel+".TenantId(listParams.TenantId))"),
		Id("return "+lowerModel+".And(wheres...)"),
	)
	err = whereF.Save("../pkg/" + pkgName + "/g_where.go")
	if err != nil {
		panic(err)
	}

	// routes
	routeF := NewFile(pkgName)

	routeF.Func().Id("RegisterRoutes").Params(Id("g *gin.RouterGroup")).Block(
		Id(`g.POST("/`+pkgName+`", `+funcAdd+`)`),
		Id(`g.GET("/`+pkgName+`", `+funcList+`)`),
		Id(`g.DELETE("/`+pkgName+`/:id", `+funcDelete+`)`),
		Id(`g.PUT("/`+pkgName+`/:id", `+funcUpdate+`)`),
	)

	err = routeF.Save("../pkg/" + pkgName + "/g_routes.go")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", routeF)
}
