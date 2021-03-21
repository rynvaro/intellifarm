package inspections

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/inspection"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func InspectionAddHandler(c *gin.Context) {
	form := &ent.Inspection{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	inspection, err := db.Client.Inspection.Create().
		SetById(form.ById).
		SetByName(form.ByName).
		SetDate(form.Date).
		SetEarNumber(form.EarNumber).
		SetHandleId(form.HandleId).
		SetHandleName(form.HandleName).
		SetItemId(form.ItemId).
		SetItemName(form.ItemName).
		SetMethodId(form.MethodId).
		SetMethodName(form.MethodName).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetResultId(form.ResultId).
		SetResultName(form.ResultName).
		SetCattleId(form.CattleId).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTenantId(form.TenantId).SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if _, err := db.Client.Event.Create().SetCreatedAt(time.Now().UnixNano()).
		SetDeleted(0).SetEarNumber(form.EarNumber).SetEventName("检疫").
		SetEventType("兽医事件").SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).Save(c.Request.Context()); err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp.Success(inspection))
}

func InspectionListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.Inspection.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	inspections, err := db.Client.Inspection.Query().Where(where).Order(ent.Desc(inspection.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   inspections,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func InspectionDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.Inspection.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func InspectionUpdateHandler(c *gin.Context) {
	form := &ent.Inspection{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	inspection, err := db.Client.Inspection.UpdateOneID(form.ID).
		SetById(form.ById).
		SetByName(form.ByName).
		SetDate(form.Date).
		SetEarNumber(form.EarNumber).
		SetHandleId(form.HandleId).
		SetHandleName(form.HandleName).
		SetItemId(form.ItemId).
		SetItemName(form.ItemName).
		SetMethodId(form.MethodId).
		SetMethodName(form.MethodName).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetResultId(form.ResultId).
		SetResultName(form.ResultName).
		SetCattleId(form.CattleId).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTenantId(form.TenantId).SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(inspection))
}
