package tenants

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/tenant"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func TenantAddHandler(c *gin.Context) {
	form := &ent.Tenant{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	tenant, err := db.Client.Tenant.Create().
		SetAddress(form.Address).
		SetCode(form.Code).
		SetCompany(form.Company).
		SetEnabled(form.Enabled).
		SetName(form.Name).
		SetPhone(form.Phone).
		SetRegion(form.Region).
		SetRemarks(form.Remarks).
		SetUserName(form.UserName).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp.Success(tenant))
}

func TenantListHandler(c *gin.Context) {
	level := c.MustGet("level").(int)
	if level != 1 {
		c.JSON(http.StatusOK, resp.Success(paging.PageData{}))
		return
	}
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	where := Where(listParams)
	totalCount, err := db.Client.Tenant.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	tenants, err := db.Client.Tenant.Query().Where(where).Order(ent.Desc(tenant.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   tenants,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func TenantDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.Tenant.DeleteOneID(int(id.Id)).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func TenantUpdateHandler(c *gin.Context) {
	form := &ent.Tenant{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	tenant, err := db.Client.Tenant.UpdateOneID(form.ID).
		SetAddress(form.Address).
		SetCode(form.Code).
		SetCompany(form.Company).
		SetEnabled(form.Enabled).
		SetName(form.Name).
		SetPhone(form.Phone).
		SetRegion(form.Region).
		SetRemarks(form.Remarks).
		SetUserName(form.UserName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(tenant))
}
