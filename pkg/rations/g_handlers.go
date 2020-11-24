package rations

import (
	"cattleai/apicommon"
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/ration"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func RationAddHandler(c *gin.Context) {
	form := &ent.Ration{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	ration, err := db.Client.Ration.Create().
		SetAdjustDate(form.AdjustDate).
		SetCode(form.Code).
		SetCost(form.Cost).
		SetCreateDate(form.CreateDate).
		SetDisableDate(form.DisableDate).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetStatus(form.Status).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, apicommon.SuccessResponse(ration))
}

func RationListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	where := Where(listParams)
	totalCount, err := db.Client.Ration.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	rations, err := db.Client.Ration.Query().Where(where).Order(ent.Desc(ration.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   rations,
		Paging: page,
	}
	c.JSON(http.StatusOK, apicommon.SuccessResponse(pageData))
}

func RationDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	ration, err := db.Client.Ration.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, apicommon.SuccessResponse(ration))
}

func RationUpdateHandler(c *gin.Context) {
	form := &ent.Ration{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	ration, err := db.Client.Ration.UpdateOneID(form.ID).
		SetAdjustDate(form.AdjustDate).
		SetCode(form.Code).
		SetCost(form.Cost).
		SetCreateDate(form.CreateDate).
		SetDisableDate(form.DisableDate).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetStatus(form.Status).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, apicommon.SuccessResponse(ration))
}
