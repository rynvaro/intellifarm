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

	tx, err := db.Client.Tx(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	//更新物料库存
	count := form.Count
	if form.Type == 2 {
		count = -count
	}
	m, err := tx.Material.UpdateOneID(int(form.MaterialId)).AddInventory(int64(count)).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	var after int64 = m.Inventory
	var before int64 = 0
	if form.Type == 2 {
		before = m.Inventory + int64(count)
	} else {
		before = m.Inventory - int64(count)
	}

	log.Debug().Msg(form.String())
	inventoryflow, err := tx.InventoryFlow.Create().
		SetCount(form.Count).
		SetDate(form.Date).
		SetStatus(form.Status).
		SetMaterialCode(form.MaterialCode).
		SetMaterialId(form.MaterialId).
		SetMaterialName(form.MaterialName).
		SetRemarks(form.Remarks).
		SetSeqNumber(form.MaterialCode + "-" + time.Now().Format("20060102150405")).
		SetType(form.Type).
		SetUnit(form.Unit).
		SetUserName(form.UserName).
		SetBefore(before).SetAfter(after).SetSysMaterialId(form.SysMaterialId).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetIsChecked(false).SetReportFileAddress(form.ReportFileAddress).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	if err := tx.Commit(); err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	c.JSON(http.StatusOK, resp.Success(inventoryflow))
}

func InventoryFlowListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
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

	flow, err := db.Client.InventoryFlow.Get(c.Request.Context(), int(id.Id))
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	count := flow.Count
	if flow.Type == 1 {
		count = -count
	}
	tx, err := db.Client.Tx(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	_, err = tx.Material.UpdateOneID(int(flow.MaterialId)).AddInventory(int64(count)).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	err = tx.InventoryFlow.DeleteOneID(int(id.Id)).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	if err := tx.Commit(); err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
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
		SetMaterialId(form.MaterialId).
		SetMaterialName(form.MaterialName).
		SetRemarks(form.Remarks).
		SetSeqNumber(form.SeqNumber).
		SetStatus(form.Status).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetType(form.Type).
		SetUnit(form.Unit).
		SetBefore(form.Before).SetAfter(form.After).SetSysMaterialId(form.SysMaterialId).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetIsChecked(false).SetReportFileAddress(form.ReportFileAddress).
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
