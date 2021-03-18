package abortions

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/abortion"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func AbortionAddHandler(c *gin.Context) {
	form := &ent.Abortion{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	abortion, err := db.Client.Abortion.Create().
		SetAbortionAt(form.AbortionAt).
		SetAbortionTypeId(form.AbortionTypeId).
		SetAbortionTypeName(form.AbortionTypeName).
		SetEarNumber(form.EarNumber).
		SetName(form.Name).
		SetPregnantAt(form.PregnantAt).
		SetRemarks(form.Remarks).
		SetReproductiveState(form.ReproductiveState).
		SetShedName(form.ShedName).
		SetTimes(form.Times).
		SetUserName(form.UserName).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if _, err := db.Client.Event.Create().SetCreatedAt(time.Now().UnixNano()).
		SetDeleted(0).SetEarNumber(form.EarNumber).SetEventName("流产").
		SetEventType("繁殖事件").SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).Save(c.Request.Context()); err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp.Success(abortion))
}

func AbortionListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.Abortion.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	abortions, err := db.Client.Abortion.Query().Where(where).Order(ent.Desc(abortion.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   abortions,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func AbortionDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.Abortion.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func AbortionUpdateHandler(c *gin.Context) {
	form := &ent.Abortion{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	abortion, err := db.Client.Abortion.UpdateOneID(form.ID).
		SetAbortionAt(form.AbortionAt).
		SetAbortionTypeId(form.AbortionTypeId).
		SetAbortionTypeName(form.AbortionTypeName).
		SetEarNumber(form.EarNumber).
		SetName(form.Name).
		SetPregnantAt(form.PregnantAt).
		SetRemarks(form.Remarks).
		SetReproductiveState(form.ReproductiveState).
		SetShedName(form.ShedName).
		SetTimes(form.Times).
		SetUserName(form.UserName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(abortion))
}
