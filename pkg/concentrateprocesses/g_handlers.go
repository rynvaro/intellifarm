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

	// concentrate, err := db.Client.Concentrate.Get(c.Request.Context(), form.ConcentrateId)
	// if err != nil {
	// 	log.Error().Msg(err.Error())
	// 	c.Status(http.StatusInternalServerError)
	// 	return
	// }

	// TODO 扣减物料库存
	// formula, err := db.Client.ConcentrateFormula.Get(c.Request.Context(), concentrate.FormulaId)
	// if err != nil {
	// 	log.Error().Msg(err.Error())
	// 	c.Status(http.StatusInternalServerError)
	// 	return
	// }

	// formulaData := []*formulaData{}

	// if err := json.Unmarshal([]byte(formula.Data), &formulaData); err != nil {
	// 	log.Error().Msg(err.Error())
	// 	c.Status(http.StatusInternalServerError)
	// 	return
	// }

	// if len(formulaData) == 0 {
	// 	log.Error().Msg(err.Error())
	// 	c.Status(http.StatusInternalServerError)
	// 	return
	// }

	tx, err := db.Client.Tx(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	_, err = tx.Concentrate.UpdateOneID(int(form.ConcentrateId)).AddInventory(form.In).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	log.Debug().Msg(form.String())
	concentrateprocess, err := tx.ConcentrateProcess.Create().
		SetConcentrateId(form.ConcentrateId).
		SetCount(form.Count).
		SetDate(form.Date).
		SetIn(form.In).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetUserName(form.UserName).
		SetTenantId(form.TenantId).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetTenantName(form.TenantName).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		tx.Rollback()
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		log.Error().Msg(err.Error())
		tx.Rollback()
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp.Success(concentrateprocess))
}

type formulaData struct {
	Mid   int64   `json:"mid"`
	Ratio float32 `json:"ratio"`
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
	err := db.Client.ConcentrateProcess.DeleteOneID(int(id.Id)).Exec(c.Request.Context())
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
		SetConcentrateId(form.ConcentrateId).
		SetCount(form.Count).
		SetDate(form.Date).
		SetIn(form.In).
		SetName(form.Name).
		SetTenantId(form.TenantId).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetTenantName(form.TenantName).
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
