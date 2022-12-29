package materialtests

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/materialtest"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func MaterialTestAddHandler(c *gin.Context) {
	form := &ent.MaterialTest{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	materialtest, err := db.Client.MaterialTest.Create().
		SetAddSeqNumber(form.AddSeqNumber).
		SetCategory(form.Category).
		SetCode(form.Code).
		SetDate(form.Date).
		SetMaterialCategory(form.MaterialCategory).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetSeqNumber(form.SeqNumber).
		SetType(form.Type).
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
	c.JSON(http.StatusOK, resp.Success(materialtest))
}

func MaterialTestListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.MaterialTest.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	materialtests, err := db.Client.MaterialTest.Query().Where(where).Order(ent.Desc(materialtest.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   materialtests,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func MaterialTestDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.MaterialTest.DeleteOneID(int(id.Id)).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func MaterialTestUpdateHandler(c *gin.Context) {
	form := &ent.MaterialTest{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	materialtest, err := db.Client.MaterialTest.UpdateOneID(form.ID).
		SetAddSeqNumber(form.AddSeqNumber).
		SetCategory(form.Category).
		SetCode(form.Code).
		SetDate(form.Date).
		SetMaterialCategory(form.MaterialCategory).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetSeqNumber(form.SeqNumber).
		SetType(form.Type).
		SetUserName(form.UserName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(materialtest))
}
