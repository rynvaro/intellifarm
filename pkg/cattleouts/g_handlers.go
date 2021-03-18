package cattleouts

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/cattleout"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CattleOutAddHandler(c *gin.Context) {
	form := &ent.CattleOut{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattleout, err := db.Client.CattleOut.Create().
		SetCheckCode(form.CheckCode).
		SetCost(form.Cost).
		SetDate(form.Date).
		SetOutType(form.OutType).
		SetRemarks(form.Remarks).
		SetShippingCode(form.ShippingCode).
		SetShippingFee(form.ShippingFee).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetTo(form.To).
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
	c.JSON(http.StatusOK, resp.Success(cattleout))
}

func CattleOutListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.CattleOut.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	cattleouts, err := db.Client.CattleOut.Query().Where(where).Order(ent.Desc(cattleout.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   cattleouts,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func CattleOutDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.CattleOut.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func CattleOutUpdateHandler(c *gin.Context) {
	form := &ent.CattleOut{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattleout, err := db.Client.CattleOut.UpdateOneID(form.ID).
		SetCheckCode(form.CheckCode).
		SetCost(form.Cost).
		SetDate(form.Date).
		SetOutType(form.OutType).
		SetRemarks(form.Remarks).
		SetShippingCode(form.ShippingCode).
		SetShippingFee(form.ShippingFee).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetTo(form.To).
		SetUserName(form.UserName).
		SetWeight(form.Weight).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattleout))
}
