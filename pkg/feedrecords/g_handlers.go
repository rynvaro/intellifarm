package feedrecords

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/feedrecord"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func FeedRecordAddHandler(c *gin.Context) {
	form := &ent.FeedRecord{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	feedrecord, err := db.Client.FeedRecord.Create().
		SetCount(form.Count).
		SetDate(form.Date).
		SetName(form.Name).
		SetRationAmount(form.RationAmount).
		SetRationCode(form.RationCode).
		SetRationName(form.RationName).
		SetRemarks(form.Remarks).
		SetShedName(form.ShedName).
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
	c.JSON(http.StatusOK, resp.Success(feedrecord))
}

func FeedRecordListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.FeedRecord.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	feedrecords, err := db.Client.FeedRecord.Query().Where(where).Order(ent.Desc(feedrecord.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   feedrecords,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func FeedRecordDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.FeedRecord.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func FeedRecordUpdateHandler(c *gin.Context) {
	form := &ent.FeedRecord{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	feedrecord, err := db.Client.FeedRecord.UpdateOneID(form.ID).
		SetCount(form.Count).
		SetDate(form.Date).
		SetName(form.Name).
		SetRationAmount(form.RationAmount).
		SetRationCode(form.RationCode).
		SetRationName(form.RationName).
		SetRemarks(form.Remarks).
		SetShedName(form.ShedName).
		SetUserName(form.UserName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(feedrecord))
}
