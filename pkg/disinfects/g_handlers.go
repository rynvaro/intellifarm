package disinfects

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/disinfect"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func DisinfectAddHandler(c *gin.Context) {
	form := &ent.Disinfect{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	disinfect, err := db.Client.Disinfect.Create().
		SetDate(form.Date).
		SetDrug(form.Drug).
		SetMethodId(form.MethodId).
		SetMethodName(form.MethodName).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetTypeId(form.TypeId).
		SetTypeName(form.TypeName).
		SetWayId(form.WayId).
		SetWayName(form.WayName).
		SetCattleId(form.CattleId).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTenantId(form.TenantId).SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(disinfect))
}

func DisinfectListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.Disinfect.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	disinfects, err := db.Client.Disinfect.Query().Where(where).Order(ent.Desc(disinfect.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   disinfects,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func DisinfectDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.Disinfect.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func DisinfectUpdateHandler(c *gin.Context) {
	form := &ent.Disinfect{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	disinfect, err := db.Client.Disinfect.UpdateOneID(form.ID).
		SetDate(form.Date).
		SetDrug(form.Drug).
		SetMethodId(form.MethodId).
		SetMethodName(form.MethodName).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetTypeId(form.TypeId).
		SetTypeName(form.TypeName).
		SetWayId(form.WayId).
		SetWayName(form.WayName).
		SetCattleId(form.CattleId).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTenantId(form.TenantId).SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(disinfect))
}
