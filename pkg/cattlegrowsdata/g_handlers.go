package cattlegrowsdata

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/cattlegrowsdata"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CattleGrowsDataAddHandler(c *gin.Context) {
	form := &ent.CattleGrowsData{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattlegrowsdata, err := db.Client.CattleGrowsData.Create().
		SetBust(form.Bust).
		SetEarNumber(form.EarNumber).
		SetHeight(form.Height).
		SetMeasuredAt(form.MeasuredAt).
		SetMeasuredBy(form.MeasuredBy).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetShedName(form.ShedName).
		SetWeight(form.Weight).
		SetTenantId(c.MustGet("tenantId").(int64)).
		SetTenantName(c.MustGet("tenantName").(string)).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if _, err := db.Client.Event.Create().SetCreatedAt(time.Now().UnixNano()).
		SetDeleted(0).SetEarNumber(form.EarNumber).SetEventName("体测").
		SetEventType("其他事件").SetTenantId(c.MustGet("tenantId").(int64)).
		SetTenantName(c.MustGet("tenantName").(string)).Save(c.Request.Context()); err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp.Success(cattlegrowsdata))
}

func CattleGrowsDataListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.TenantId = c.MustGet("tenantId").(int64)
	where := Where(listParams)
	totalCount, err := db.Client.CattleGrowsData.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	cattlegrowsdata, err := db.Client.CattleGrowsData.Query().Where(where).Order(ent.Desc(cattlegrowsdata.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   cattlegrowsdata,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func CattleGrowsDataDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	cattlegrowsdata, err := db.Client.CattleGrowsData.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattlegrowsdata))
}

func CattleGrowsDataUpdateHandler(c *gin.Context) {
	form := &ent.CattleGrowsData{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattlegrowsdata, err := db.Client.CattleGrowsData.UpdateOneID(form.ID).
		SetBust(form.Bust).
		SetEarNumber(form.EarNumber).
		SetHeight(form.Height).
		SetMeasuredAt(form.MeasuredAt).
		SetMeasuredBy(form.MeasuredBy).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetShedName(form.ShedName).
		SetWeight(form.Weight).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattlegrowsdata))
}
