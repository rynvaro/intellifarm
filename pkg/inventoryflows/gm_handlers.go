package inventoryflows

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/inventoryflow"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func InventoryFlowAddHandler(c *gin.Context) {
	form := &ent.InventoryFlow{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	inventoryflow, err := db.Client.InventoryFlow.Create().
		SetCount(form.Count).
		SetDate(form.Date).
		SetStatus(form.Status).
		SetMaterialCode(form.MaterialCode).
		SetMaterialID(form.MaterialID).
		SetMaterialName(form.MaterialName).
		SetRemarks(form.Remarks).
		SetSeqNumber(form.MaterialCode + "-" + time.Now().Format("20060102150405") + "-" + strconv.FormatInt(time.Now().UnixNano(), 10)).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetType(form.Type).
		SetUnit(form.Unit).
		SetUserName(form.UserName).
		SetTenantId(c.MustGet("tenantId").(int64)).
		SetTenantName(c.MustGet("tenantName").(string)).SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	//更新物料库存
	count := form.Count
	if form.Type == 2 {
		count = -count
	}
	m, err := db.Client.Material.UpdateOneID(form.MaterialID).AddInventory(int64(count)).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	log.Info().Msg(m.String())
	c.JSON(http.StatusOK, resp.Success(inventoryflow))
}

func InventoryFlowListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.TenantId = c.MustGet("tenantId").(int64)
	where := Where(listParams)
	totalCount, err := db.Client.InventoryFlow.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	inventoryflows, err := db.Client.InventoryFlow.Query().Where(where).Order(ent.Desc(inventoryflow.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   inventoryflows,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func InventoryFlowDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	inventoryflow, err := db.Client.InventoryFlow.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(inventoryflow))
}

func InventoryFlowUpdateHandler(c *gin.Context) {
	form := &ent.InventoryFlow{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	inventoryflow, err := db.Client.InventoryFlow.UpdateOneID(form.ID).
		SetCount(form.Count).
		SetDate(form.Date).
		SetMaterialCode(form.MaterialCode).
		SetMaterialID(form.MaterialID).
		SetMaterialName(form.MaterialName).
		SetRemarks(form.Remarks).
		SetSeqNumber(form.SeqNumber).
		SetStatus(form.Status).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetType(form.Type).
		SetUnit(form.Unit).
		SetUserName(form.UserName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(inventoryflow))
}
