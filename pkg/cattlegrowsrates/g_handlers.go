package cattlegrowsrates

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/cattlegrowsrate"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CattleGrowsRateAddHandler(c *gin.Context) {
	form := &ent.CattleGrowsRate{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattlegrowsrate, err := db.Client.CattleGrowsRate.Create().
		SetEarNumber(form.EarNumber).
		SetName(form.Name).
		SetRate(form.Rate).
		SetRatedAt(form.RatedAt).
		SetRatedBy(form.RatedBy).
		SetRemarks(form.Remarks).
		SetShedName(form.ShedName).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	// if _, err := db.Client.Event.Create().SetCreatedAt(time.Now().UnixNano()).
	// 	SetDeleted(0).SetEarNumber(form.EarNumber).SetEventTypeName("体况评分").
	// 	SetEventSubTypeName("其他事件").SetTenantId(form.TenantId).
	// 	SetFarmId(form.FarmId).SetFarmName(form.FarmName).
	// 	SetShedId(form.ShedId).SetShedName(form.ShedName).
	// 	SetTenantName(form.TenantName).Save(c.Request.Context()); err != nil {
	// 	log.Error().Msg(err.Error())
	// 	c.Status(http.StatusInternalServerError)
	// 	return
	// }

	c.JSON(http.StatusOK, resp.Success(cattlegrowsrate))
}

func CattleGrowsRateListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.CattleGrowsRate.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	cattlegrowsrates, err := db.Client.CattleGrowsRate.Query().Where(where).Order(ent.Desc(cattlegrowsrate.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   cattlegrowsrates,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func CattleGrowsRateDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.CattleGrowsRate.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func CattleGrowsRateUpdateHandler(c *gin.Context) {
	form := &ent.CattleGrowsRate{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattlegrowsrate, err := db.Client.CattleGrowsRate.UpdateOneID(form.ID).
		SetEarNumber(form.EarNumber).
		SetName(form.Name).
		SetRate(form.Rate).
		SetRatedAt(form.RatedAt).
		SetRatedBy(form.RatedBy).
		SetRemarks(form.Remarks).
		SetShedName(form.ShedName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattlegrowsrate))
}
