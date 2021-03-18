package cattledies

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/cattledie"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CattleDieAddHandler(c *gin.Context) {
	form := &ent.CattleDie{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattledie, err := db.Client.CattleDie.Create().
		SetDUserName(form.DUserName).
		SetDate(form.Date).
		SetDeclared(form.Declared).
		SetEarNumber(form.EarNumber).
		SetHandleMethod(form.HandleMethod).
		SetInsured(form.Insured).
		SetInsuredCode(form.InsuredCode).
		SetInsuredCompany(form.InsuredCompany).
		SetReason(form.Reason).
		SetRemarks(form.Remarks).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetUserName(form.UserName).
		SetWeight(form.Weight).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattledie))
}

func CattleDieListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.CattleDie.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	cattledies, err := db.Client.CattleDie.Query().Where(where).Order(ent.Desc(cattledie.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   cattledies,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func CattleDieDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.CattleDie.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func CattleDieUpdateHandler(c *gin.Context) {
	form := &ent.CattleDie{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattledie, err := db.Client.CattleDie.UpdateOneID(form.ID).
		SetDUserName(form.DUserName).
		SetDate(form.Date).
		SetDeclared(form.Declared).
		SetEarNumber(form.EarNumber).
		SetHandleMethod(form.HandleMethod).
		SetInsured(form.Insured).
		SetInsuredCode(form.InsuredCode).
		SetInsuredCompany(form.InsuredCompany).
		SetReason(form.Reason).
		SetRemarks(form.Remarks).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetUserName(form.UserName).
		SetWeight(form.Weight).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattledie))
}
