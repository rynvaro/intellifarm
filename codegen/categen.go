package main

import (
	"fmt"
	"os"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/gertd/go-pluralize"
	"github.com/rs/zerolog/log"
)

var (
	ErrStatement = If(Err().Op("!=").Id("nil")).Block(
		Id("log.Error().Msg(err.Error())"),
		Id("return"),
	)
	ErrStatementWithResp = If(Err().Op("!=").Id("nil")).Block(
		Id("log.Error().Msg(err.Error())"),
		Id("c.Status(http.StatusInternalServerError)"),
		Id("return"),
	)
)

// 生成各种类别字典数据代码
func GenCates(model string) {
	if model == "" {
		panic("empty")
	}
	pluralize := pluralize.NewClient()

	param := "c *gin.Context"
	modelPlural := pluralize.Plural(model)
	modelPluralLowerCase := strings.ToLower(modelPlural)

	pkgName := modelPluralLowerCase

	f := NewFile(pkgName)

	f.Func().Id(modelPlural).Params(Id(param)).Block(
		List(Id(modelPluralLowerCase), Err()).Op(":=").Id("db.Client."+model+".Query().All(c.Request.Context())"),
		ErrStatementWithResp,
		Id("c.JSON(http.StatusOK, resp.Success("+modelPluralLowerCase+"))"),
	).Line()

	if err := os.Mkdir("../pkg/"+pkgName, 0755); err != nil {
		log.Warn().Msg(err.Error())
	}

	err := f.Save("../pkg/" + pkgName + "/g_handlers.go")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", f)

	// routes
	routeF := NewFile(pkgName)

	routeF.Func().Id("RegisterRoutes").Params(Id("g *gin.RouterGroup")).Block(
		Id(`g.GET("/` + pkgName + `", ` + modelPlural + `)`),
	)

	err = routeF.Save("../pkg/" + pkgName + "/g_routes.go")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", routeF)
}
