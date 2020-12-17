package immunities

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/immunity"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ImmunityAddHandler(c *gin.Context) {
	form := &ent.Immunity{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	immunity, err := db.Client.Immunity.Create().
		SetDate(form.Date).
		SetDrug(form.Drug).
		SetEarNumber(form.EarNumber).
		SetItemId(form.ItemId).
		SetItemName(form.ItemName).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetShedName(form.ShedName).
		SetUserName(form.UserName).
		SetTenantId(c.MustGet("tenantId").(int64)).
		SetTenantName(c.MustGet("tenantName").(string)).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(immunity))
}

func ImmunityListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.TenantId = c.MustGet("tenantId").(int64)
	where := Where(listParams)
	totalCount, err := db.Client.Immunity.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	immunities, err := db.Client.Immunity.Query().Where(where).Order(ent.Desc(immunity.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   immunities,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func ImmunityDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	immunity, err := db.Client.Immunity.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(immunity))
}

func ImmunityUpdateHandler(c *gin.Context) {
	form := &ent.Immunity{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	immunity, err := db.Client.Immunity.UpdateOneID(form.ID).
		SetDate(form.Date).
		SetDrug(form.Drug).
		SetEarNumber(form.EarNumber).
		SetItemId(form.ItemId).
		SetItemName(form.ItemName).
		SetName(form.Name).
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
	c.JSON(http.StatusOK, resp.Success(immunity))
}
