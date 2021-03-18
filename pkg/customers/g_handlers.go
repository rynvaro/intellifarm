package customers

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/customer"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CustomerAddHandler(c *gin.Context) {
	form := &ent.Customer{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	customer, err := db.Client.Customer.Create().
		SetAccount(form.Account).
		SetAddress(form.Address).
		SetBank(form.Bank).
		SetEmail(form.Email).
		SetGender(form.Gender).
		SetName(form.Name).
		SetPhone(form.Phone).
		SetProvince(form.Province).
		SetRemarks(form.Remarks).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetType(form.Type).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(customer))
}

func CustomerListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.Customer.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	customers, err := db.Client.Customer.Query().Where(where).Order(ent.Desc(customer.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   customers,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func CustomerDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.Customer.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func CustomerUpdateHandler(c *gin.Context) {
	form := &ent.Customer{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	customer, err := db.Client.Customer.UpdateOneID(form.ID).
		SetAccount(form.Account).
		SetAddress(form.Address).
		SetBank(form.Bank).
		SetEmail(form.Email).
		SetGender(form.Gender).
		SetName(form.Name).
		SetPhone(form.Phone).
		SetProvince(form.Province).
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
	c.JSON(http.StatusOK, resp.Success(customer))
}
