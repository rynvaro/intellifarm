package concentrateprocesses

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/concentrateprocess"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ConcentrateProcessAddHandler(c *gin.Context) {
	form := &ent.ConcentrateProcess{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	concentrateprocess, err := db.Client.ConcentrateProcess.Create().
		SetCode(form.Code).
		SetFormulaID(form.FormulaID).
		SetCount(form.Count).
		SetDate(form.Date).
		SetIn(form.In).
		SetInventory(form.Inventory).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetUserName(form.UserName).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(concentrateprocess))
}

func ConcentrateProcessListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.ConcentrateProcess.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	concentrateprocesses, err := db.Client.ConcentrateProcess.Query().Where(where).Order(ent.Desc(concentrateprocess.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   concentrateprocesses,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func ConcentrateProcessDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.ConcentrateProcess.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func ConcentrateProcessUpdateHandler(c *gin.Context) {
	form := &ent.ConcentrateProcess{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	concentrateprocess, err := db.Client.ConcentrateProcess.UpdateOneID(form.ID).
		SetCode(form.Code).
		SetCount(form.Count).
		SetFormulaID(form.FormulaID).
		SetDate(form.Date).
		SetIn(form.In).
		SetInventory(form.Inventory).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetUserName(form.UserName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(concentrateprocess))
}
