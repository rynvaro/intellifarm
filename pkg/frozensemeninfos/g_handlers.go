package frozensemeninfos

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/frozensemeninfo"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func FrozenSemenInfoAddHandler(c *gin.Context) {
	form := &ent.FrozenSemenInfo{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	frozensemeninfo, err := db.Client.FrozenSemenInfo.Create().
		SetBirthday(form.Birthday).
		SetBullNumber(form.BullNumber).
		SetCode(form.Code).
		SetFrom(form.From).
		SetName(form.Name).
		SetRegCode(form.RegCode).
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
	c.JSON(http.StatusOK, resp.Success(frozensemeninfo))
}

func FrozenSemenInfoListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.TenantId = c.MustGet("tenantId").(int64)
	where := Where(listParams)
	totalCount, err := db.Client.FrozenSemenInfo.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	frozensemeninfos, err := db.Client.FrozenSemenInfo.Query().Where(where).Order(ent.Desc(frozensemeninfo.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   frozensemeninfos,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func FrozenSemenInfoDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	frozensemeninfo, err := db.Client.FrozenSemenInfo.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(frozensemeninfo))
}

func FrozenSemenInfoUpdateHandler(c *gin.Context) {
	form := &ent.FrozenSemenInfo{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	frozensemeninfo, err := db.Client.FrozenSemenInfo.UpdateOneID(form.ID).
		SetBirthday(form.Birthday).
		SetBullNumber(form.BullNumber).
		SetCode(form.Code).
		SetFrom(form.From).
		SetName(form.Name).
		SetRegCode(form.RegCode).
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
	c.JSON(http.StatusOK, resp.Success(frozensemeninfo))
}
