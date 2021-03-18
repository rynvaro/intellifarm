package cattlegrows

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/cattlegrow"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CattleGrowAddHandler(c *gin.Context) {
	form := &ent.CattleGrow{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattlegrow, err := db.Client.CattleGrow.Create().
		SetConversionRate(form.ConversionRate).
		SetDailyFeedWeight(form.DailyFeedWeight).
		SetDailyWeight(form.DailyWeight).
		SetDateEnd(form.DateEnd).
		SetDateStart(form.DateStart).
		SetEarNumber(form.EarNumber).
		SetFeedWeight(form.FeedWeight).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetStage(form.Stage).
		SetUserName(form.UserName).
		SetWeightEnd(form.WeightEnd).
		SetWeightStart(form.WeightStart).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattlegrow))
}

func CattleGrowListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.CattleGrow.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	cattlegrows, err := db.Client.CattleGrow.Query().Where(where).Order(ent.Desc(cattlegrow.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   cattlegrows,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func CattleGrowDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.CattleGrow.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func CattleGrowUpdateHandler(c *gin.Context) {
	form := &ent.CattleGrow{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattlegrow, err := db.Client.CattleGrow.UpdateOneID(form.ID).
		SetConversionRate(form.ConversionRate).
		SetDailyFeedWeight(form.DailyFeedWeight).
		SetDailyWeight(form.DailyWeight).
		SetDateEnd(form.DateEnd).
		SetDateStart(form.DateStart).
		SetEarNumber(form.EarNumber).
		SetFeedWeight(form.FeedWeight).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetStage(form.Stage).
		SetUserName(form.UserName).
		SetWeightEnd(form.WeightEnd).
		SetWeightStart(form.WeightStart).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattlegrow))
}
