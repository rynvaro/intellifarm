package cattleins

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/cattlein"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CattleInAddHandler(c *gin.Context) {
	form := &ent.CattleIn{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattlein, err := db.Client.CattleIn.Create().
		SetCheckCode(form.CheckCode).
		SetCost(form.Cost).
		SetDate(form.Date).
		SetFrom(form.From).
		SetJoinType(form.JoinType).
		SetRemarks(form.Remarks).
		SetShippingCode(form.ShippingCode).
		SetShippingFee(form.ShippingFee).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetUserName(form.UserName).
		SetWeight(form.Weight).
		SetTenantId(c.MustGet("tenantId").(int64)).
		SetTenantName(c.MustGet("tenantName").(string)).SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattlein))
}

func CattleInListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.TenantId = c.MustGet("tenantId").(int64)
	where := Where(listParams)
	totalCount, err := db.Client.CattleIn.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	cattleins, err := db.Client.CattleIn.Query().Where(where).Order(ent.Desc(cattlein.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   cattleins,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func CattleInDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	cattlein, err := db.Client.CattleIn.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattlein))
}

func CattleInUpdateHandler(c *gin.Context) {
	form := &ent.CattleIn{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattlein, err := db.Client.CattleIn.UpdateOneID(form.ID).
		SetCheckCode(form.CheckCode).
		SetCost(form.Cost).
		SetDate(form.Date).
		SetFrom(form.From).
		SetJoinType(form.JoinType).
		SetRemarks(form.Remarks).
		SetShippingCode(form.ShippingCode).
		SetShippingFee(form.ShippingFee).
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
	c.JSON(http.StatusOK, resp.Success(cattlein))
}
