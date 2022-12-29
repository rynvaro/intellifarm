package concentrateformulas

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/concentrateformula"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ConcentrateFormulaAddHandler(c *gin.Context) {
	form := &ent.ConcentrateFormula{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	concentrateformula, err := db.Client.ConcentrateFormula.Create().
		SetCode(form.Code).
		SetCost(form.Cost).
		SetData(form.Data).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetStatus(form.Status).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(concentrateformula))
}

func ConcentrateFormulaListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.ConcentrateFormula.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	concentrateformulas, err := db.Client.ConcentrateFormula.Query().Where(where).Order(ent.Desc(concentrateformula.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   concentrateformulas,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func ConcentrateFormulaDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.ConcentrateFormula.DeleteOneID(int(id.Id)).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func ConcentrateFormulaUpdateHandler(c *gin.Context) {
	form := &ent.ConcentrateFormula{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	concentrateformula, err := db.Client.ConcentrateFormula.UpdateOneID(form.ID).
		SetCode(form.Code).
		SetData(form.Data).
		SetCost(form.Cost).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetStatus(form.Status).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(concentrateformula))
}
