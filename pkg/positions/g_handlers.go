package positions

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/position"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func PositionAddHandler(c *gin.Context) {
	form := &ent.Position{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	position, err := db.Client.Position.Create().
		SetCode(form.Code).
		SetName(form.Name).
		SetOrder(form.Order).
		SetRemarks(form.Remarks).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetTenantId(c.MustGet("tenantId").(int64)).
		SetTenantName(c.MustGet("tenantName").(string)).SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(position))
}

func PositionListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.TenantId = c.MustGet("tenantId").(int64)
	where := Where(listParams)
	totalCount, err := db.Client.Position.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	positions, err := db.Client.Position.Query().Where(where).Order(ent.Desc(position.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   positions,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func PositionDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	position, err := db.Client.Position.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(position))
}

func PositionUpdateHandler(c *gin.Context) {
	form := &ent.Position{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	position, err := db.Client.Position.UpdateOneID(form.ID).
		SetCode(form.Code).
		SetName(form.Name).
		SetOrder(form.Order).
		SetRemarks(form.Remarks).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(position))
}
