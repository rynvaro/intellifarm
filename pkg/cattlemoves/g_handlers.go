package cattlemoves

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/cattlemove"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CattleMoveAddHandler(c *gin.Context) {
	form := &ent.CattleMove{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattlemove, err := db.Client.CattleMove.Create().
		SetDate(form.Date).
		SetEarNumber(form.EarNumber).
		SetCattleId(form.CattleId).
		SetReasonId(form.ReasonId).
		SetReasonName(form.ReasonName).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTenantId(form.TenantId).SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetFromShed(form.FromShed).
		SetFromShedId(form.FromShedId).
		SetRemarks(form.Remarks).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetToShedId(form.ToShedId).
		SetToShed(form.ToShed).
		SetUserName(form.UserName).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	// TODO check shed exists
	if _, err := db.Client.Cattle.UpdateOneID(int(form.CattleId)).SetShedId(form.ToShedId).
		SetShedName(form.ToShed).Save(c.Request.Context()); err != nil {
		log.Warn().Msg(err.Error())
	}
	c.JSON(http.StatusOK, resp.Success(cattlemove))
}

func CattleMoveListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.CattleMove.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	cattlemoves, err := db.Client.CattleMove.Query().Where(where).Order(ent.Desc(cattlemove.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   cattlemoves,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func CattleMoveDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.CattleMove.DeleteOneID(int(id.Id)).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func CattleMoveUpdateHandler(c *gin.Context) {
	form := &ent.CattleMove{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattlemove, err := db.Client.CattleMove.UpdateOneID(form.ID).
		SetDate(form.Date).
		SetEarNumber(form.EarNumber).
		SetCattleId(form.CattleId).
		SetReasonId(form.ReasonId).
		SetReasonName(form.ReasonName).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTenantId(form.TenantId).SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetFromShed(form.FromShed).
		SetFromShedId(form.FromShedId).
		SetRemarks(form.Remarks).
		SetToShedId(form.ToShedId).
		SetToShed(form.ToShed).
		SetUserName(form.UserName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattlemove))
}
