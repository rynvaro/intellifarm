package warehousesettings

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/warehousesetting"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func WarehouseSettingAddHandler(c *gin.Context) {
	form := &ent.WarehouseSetting{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	warehousesetting, err := db.Client.WarehouseSetting.Create().
		SetName(form.Name).
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
	c.JSON(http.StatusOK, resp.Success(warehousesetting))
}

func WarehouseSettingListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.TenantId = c.MustGet("tenantId").(int64)
	where := Where(listParams)
	totalCount, err := db.Client.WarehouseSetting.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	warehousesettings, err := db.Client.WarehouseSetting.Query().Where(where).Order(ent.Desc(warehousesetting.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   warehousesettings,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func WarehouseSettingDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	warehousesetting, err := db.Client.WarehouseSetting.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(warehousesetting))
}

func WarehouseSettingUpdateHandler(c *gin.Context) {
	form := &ent.WarehouseSetting{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	warehousesetting, err := db.Client.WarehouseSetting.UpdateOneID(form.ID).
		SetName(form.Name).
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
	c.JSON(http.StatusOK, resp.Success(warehousesetting))
}
