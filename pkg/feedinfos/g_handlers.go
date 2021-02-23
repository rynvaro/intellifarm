package feedinfos

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/feedinfo"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func FeedInfoAddHandler(c *gin.Context) {
	form := &ent.FeedInfo{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	feedinfo, err := db.Client.FeedInfo.Create().
		SetCode(form.Code).
		SetDescription(form.Description).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetType(form.Type).
		SetTenantId(c.MustGet("tenantId").(int64)).
		SetTenantName(c.MustGet("tenantName").(string)).SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(feedinfo))
}

func FeedInfoListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.TenantId = c.MustGet("tenantId").(int64)
	where := Where(listParams)
	totalCount, err := db.Client.FeedInfo.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	feedinfos, err := db.Client.FeedInfo.Query().Where(where).Order(ent.Desc(feedinfo.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   feedinfos,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func FeedInfoDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	feedinfo, err := db.Client.FeedInfo.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(feedinfo))
}

func FeedInfoUpdateHandler(c *gin.Context) {
	form := &ent.FeedInfo{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	feedinfo, err := db.Client.FeedInfo.UpdateOneID(form.ID).
		SetCode(form.Code).
		SetDescription(form.Description).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetType(form.Type).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(feedinfo))
}
