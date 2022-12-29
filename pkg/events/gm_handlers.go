package events

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/event"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func EventAddHandler(c *gin.Context) {
	form := &ent.Event{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	event, err := db.Client.Event.Create().
		SetEarNumber(form.EarNumber).
		SetEventTypeName(form.EventTypeName).
		SetEventSubTypeName(form.EventSubTypeName).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).SetCreatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(event))
}

type EventSearchParams struct {
	EarNumber   int64  `json:"earNumber" form:"earNumber"`
	InGroup     int    `json:"inGroup" form:"inGroup"` // 是否在群
	TypeName    string `json:"typeName" form:"typeName"`
	SubTypeName string `json:"subTypeName" form:"subTypeName"`
	TimeStart   int64  `json:"timeStart" form:"timeStart"`
	TimeEnd     int64  `json:"timeEnd" form:"timeEnd"`
	PStart      int    `json:"pStart" form:"pStart"`
	PEnd        int    `json:"pEnd" form:"pEnd"`
	params.ListParams
}

func EventListHandler(c *gin.Context) {
	listParams := &EventSearchParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.Event.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	events, err := db.Client.Event.Query().Where(where).Order(ent.Desc(event.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   events,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func EventDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.Event.DeleteOneID(int(id.Id)).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func EventUpdateHandler(c *gin.Context) {
	form := &ent.Event{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	event, err := db.Client.Event.UpdateOneID(form.ID).
		SetEarNumber(form.EarNumber).
		SetEventTypeName(form.EventTypeName).
		SetEventSubTypeName(form.EventSubTypeName).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(event))
}
