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
		SetAdjustDate(form.AdjustDate).
		SetCode(form.Code).
		SetCost(form.Cost).
		SetCreateDate(form.CreateDate).
		SetDisableDate(form.DisableDate).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetStatus(form.Status).
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
	concentrateformula, err := db.Client.ConcentrateFormula.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(concentrateformula))
}

func ConcentrateFormulaUpdateHandler(c *gin.Context) {
	form := &ent.ConcentrateFormula{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	concentrateformula, err := db.Client.ConcentrateFormula.UpdateOneID(form.ID).
		SetAdjustDate(form.AdjustDate).
		SetCode(form.Code).
		SetCost(form.Cost).
		SetCreateDate(form.CreateDate).
		SetDisableDate(form.DisableDate).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetStatus(form.Status).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(concentrateformula))
}
