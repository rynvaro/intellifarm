package rationprocess

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/rationprocess"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func RationProcessAddHandler(c *gin.Context) {
	form := &ent.RationProcess{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}

	// ration, err := db.Client.Ration.Get(c.Request.Context(), form.RationId)
	// if err != nil {
	// 	log.Error().Msg(err.Error())
	// 	c.Status(http.StatusInternalServerError)
	// 	return
	// }

	// TODO 扣减物料库存
	// formula, err := db.Client.RationFormula.Get(c.Request.Context(), ration.FormulaId)
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

	_, err = tx.Ration.UpdateOneID(form.RationId).AddInventory(form.In).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	log.Debug().Msg(form.String())
	rationprocess, err := tx.RationProcess.Create().
		SetRationId(form.RationId).
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

	c.JSON(http.StatusOK, resp.Success(rationprocess))
}

type formulaData struct {
	Mid   int64   `json:"mid"`
	Ratio float32 `json:"ratio"`
}

func RationProcessListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.RationProcess.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	rationprocess, err := db.Client.RationProcess.Query().Where(where).Order(ent.Desc(rationprocess.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   rationprocess,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func RationProcessDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.RationProcess.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func RationProcessUpdateHandler(c *gin.Context) {
	form := &ent.RationProcess{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	rationprocess, err := db.Client.RationProcess.UpdateOneID(form.ID).
		SetRationId(form.RationId).
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
	c.JSON(http.StatusOK, resp.Success(rationprocess))
}
