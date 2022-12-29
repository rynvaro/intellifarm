package epidemics

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/epidemic"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func EpidemicAddHandler(c *gin.Context) {
	form := &ent.Epidemic{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	epidemic, err := db.Client.Epidemic.Create().
		SetDiagedBy(form.DiagedBy).
		SetEarNumber(form.EarNumber).
		SetEpidemicTypeId(form.EpidemicTypeId).
		SetEpidemicTypeName(form.EpidemicTypeName).
		SetIsolatedShedName(form.IsolatedShedName).
		SetName(form.Name).
		SetTreatmentState(form.TreatmentState).
		SetOnset(form.Onset).
		SetRemarks(form.Remarks).
		SetTreatmentAt(form.TreatmentAt).
		SetTreatmentResultId(form.TreatmentResultId).
		SetTreatmentResultName(form.TreatmentResultName).
		SetWhereabout(form.Whereabout).
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
		SetDeleted(0).SetEarNumber(form.EarNumber).SetEventTypeName("发病").
		SetEventSubTypeName("兽医事件").SetTenantId(form.TenantId).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTenantName(form.TenantName).Save(c.Request.Context()); err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp.Success(epidemic))
}

func EpidemicListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.Epidemic.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	epidemics, err := db.Client.Epidemic.Query().Where(where).Order(ent.Desc(epidemic.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   epidemics,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func EpidemicDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.Epidemic.DeleteOneID(int(id.Id)).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func EpidemicUpdateHandler(c *gin.Context) {
	form := &ent.Epidemic{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	epidemic, err := db.Client.Epidemic.UpdateOneID(form.ID).
		SetDiagedBy(form.DiagedBy).
		SetEarNumber(form.EarNumber).
		SetEpidemicTypeId(form.EpidemicTypeId).
		SetEpidemicTypeName(form.EpidemicTypeName).
		SetIsolatedShedName(form.IsolatedShedName).
		SetName(form.Name).
		SetTreatmentState(form.TreatmentState).
		SetOnset(form.Onset).
		SetRemarks(form.Remarks).
		SetCattleId(form.CattleId).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTenantId(form.TenantId).SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetTreatmentAt(form.TreatmentAt).
		SetTreatmentResultId(form.TreatmentResultId).
		SetTreatmentResultName(form.TreatmentResultName).
		SetWhereabout(form.Whereabout).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(epidemic))
}
