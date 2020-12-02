package apis

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/api"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func APIAddHandler(c *gin.Context) {
	form := &ent.API{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	api, err := db.Client.API.Create().
		SetComponent(form.Component).
		SetHasSub(form.HasSub).
		SetHash(form.Hash).
		SetID(form.ID).
		SetIsSub(form.IsSub).
		SetLevel(form.Level).
		SetName(form.Name).
		SetParentId(form.ParentId).
		SetPath(form.Path).
		SetRedirect(form.Redirect).
		SetSingle(form.Single).
		SetTenantId(form.TenantId).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(api))
}

func APIListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	where := Where(listParams)
	totalCount, err := db.Client.API.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	apis, err := db.Client.API.Query().Where(where).Order(ent.Desc(api.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   apis,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func APIDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	api, err := db.Client.API.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(api))
}

func APIUpdateHandler(c *gin.Context) {
	form := &ent.API{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	api, err := db.Client.API.UpdateOneID(form.ID).
		SetComponent(form.Component).
		SetHasSub(form.HasSub).
		SetHash(form.Hash).
		SetIsSub(form.IsSub).
		SetLevel(form.Level).
		SetName(form.Name).
		SetParentId(form.ParentId).
		SetPath(form.Path).
		SetRedirect(form.Redirect).
		SetSingle(form.Single).
		SetTenantId(form.TenantId).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(api))
}
