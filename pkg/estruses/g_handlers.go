package estruses

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/estrus"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func EstrusAddHandler(c *gin.Context) {
	form := &ent.Estrus{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	estrus, err := db.Client.Estrus.Create().
		SetEarNumber(form.EarNumber).
		SetEstrusAt(form.EstrusAt).
		SetEstrusTypeId(form.EstrusTypeId).
		SetEstrusTypeName(form.EstrusTypeName).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetTimes(form.Times).
		SetUserName(form.UserName).
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

	if _, err := db.Client.Event.Create().SetCreatedAt(time.Now().UnixNano()).
		SetDeleted(0).SetEarNumber(form.EarNumber).SetEventName("发情").
		SetEventType("繁殖事件").SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).Save(c.Request.Context()); err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp.Success(estrus))
}

func EstrusListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.Estrus.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	estruses, err := db.Client.Estrus.Query().Where(where).Order(ent.Desc(estrus.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   estruses,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func EstrusDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.Estrus.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func EstrusUpdateHandler(c *gin.Context) {
	form := &ent.Estrus{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	estrus, err := db.Client.Estrus.UpdateOneID(form.ID).
		SetEarNumber(form.EarNumber).
		SetEstrusAt(form.EstrusAt).
		SetEstrusTypeId(form.EstrusTypeId).
		SetEstrusTypeName(form.EstrusTypeName).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetTimes(form.Times).
		SetCattleId(form.CattleId).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTenantId(form.TenantId).SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetUserName(form.UserName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(estrus))
}
