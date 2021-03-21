package healthcares

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/healthcare"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func HealthCareAddHandler(c *gin.Context) {
	form := &ent.HealthCare{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	healthcare, err := db.Client.HealthCare.Create().
		SetDate(form.Date).
		SetEarNumber(form.EarNumber).
		SetHoofArea(form.HoofArea).
		SetHornMethod(form.HornMethod).
		SetMethod(form.Method).
		SetReason(form.Reason).
		SetRemarks(form.Remarks).
		SetCattleId(form.CattleId).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTenantId(form.TenantId).SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetVetName(form.VetName).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if _, err := db.Client.Event.Create().SetCreatedAt(time.Now().UnixNano()).
		SetDeleted(0).SetEarNumber(form.EarNumber).SetEventName("保健").
		SetEventType("兽医事件").SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).Save(c.Request.Context()); err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if _, err := db.Client.Event.Create().SetCreatedAt(time.Now().UnixNano()).
		SetDeleted(0).SetEarNumber(form.EarNumber).SetEventName(form.Reason).
		SetEventType("其他事件").SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).Save(c.Request.Context()); err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp.Success(healthcare))
}

func HealthCareListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.HealthCare.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	healthcares, err := db.Client.HealthCare.Query().Where(where).Order(ent.Desc(healthcare.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   healthcares,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func HealthCareDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.HealthCare.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func HealthCareUpdateHandler(c *gin.Context) {
	form := &ent.HealthCare{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	healthcare, err := db.Client.HealthCare.UpdateOneID(form.ID).
		SetDate(form.Date).
		SetEarNumber(form.EarNumber).
		SetHoofArea(form.HoofArea).
		SetHornMethod(form.HornMethod).
		SetMethod(form.Method).
		SetReason(form.Reason).
		SetRemarks(form.Remarks).
		SetCattleId(form.CattleId).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTenantId(form.TenantId).SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetVetName(form.VetName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(healthcare))
}
