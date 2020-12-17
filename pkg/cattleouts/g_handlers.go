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
		SetCost(form.Cost).
		SetDate(form.Date).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetShippingFee(form.ShippingFee).
		SetTestCertificateNumber(form.TestCertificateNumber).
		SetTo(form.To).
		SetTransportCertificateNumber(form.TransportCertificateNumber).
		SetType(form.Type).
		SetUserName(form.UserName).
		SetWeight(form.Weight).
		SetTenantId(c.MustGet("tenantId").(int64)).
		SetTenantName(c.MustGet("tenantName").(string)).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
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
	listParams.TenantId = c.MustGet("tenantId").(int64)
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
	cattleout, err := db.Client.CattleOut.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattleout))
}

func CattleOutUpdateHandler(c *gin.Context) {
	form := &ent.CattleOut{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattleout, err := db.Client.CattleOut.UpdateOneID(form.ID).
		SetCost(form.Cost).
		SetDate(form.Date).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetShippingFee(form.ShippingFee).
		SetTestCertificateNumber(form.TestCertificateNumber).
		SetTo(form.To).
		SetTransportCertificateNumber(form.TransportCertificateNumber).
		SetType(form.Type).
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
