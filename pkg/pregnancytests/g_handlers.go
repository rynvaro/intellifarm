package pregnancytests

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/pregnancytest"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func PregnancyTestAddHandler(c *gin.Context) {
	form := &ent.PregnancyTest{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	pregnancytest, err := db.Client.PregnancyTest.Create().
		SetBreedingAt(form.BreedingAt).
		SetEarNumber(form.EarNumber).
		SetName(form.Name).
		SetPregnancyTestMethodId(form.PregnancyTestMethodId).
		SetPregnancyTestMethodName(form.PregnancyTestMethodName).
		SetPregnancyTestResultId(form.PregnancyTestResultId).
		SetPregnancyTestResultName(form.PregnancyTestResultName).
		SetPregnancyTestTypeId(form.PregnancyTestTypeId).
		SetPregnancyTestTypeName(form.PregnancyTestTypeName).
		SetRemarks(form.Remarks).
		SetCattleId(form.CattleId).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTenantId(form.TenantId).SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetTestAt(form.TestAt).
		SetTimes(form.Times).
		SetUserName(form.UserName).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if _, err := db.Client.Event.Create().SetCreatedAt(time.Now().UnixNano()).
		SetDeleted(0).SetEarNumber(form.EarNumber).SetEventName("孕检").
		SetEventType("繁殖事件").SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).Save(c.Request.Context()); err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp.Success(pregnancytest))
}

func PregnancyTestListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.PregnancyTest.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	pregnancytests, err := db.Client.PregnancyTest.Query().Where(where).Order(ent.Desc(pregnancytest.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   pregnancytests,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func PregnancyTestDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.PregnancyTest.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func PregnancyTestUpdateHandler(c *gin.Context) {
	form := &ent.PregnancyTest{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	pregnancytest, err := db.Client.PregnancyTest.UpdateOneID(form.ID).
		SetBreedingAt(form.BreedingAt).
		SetEarNumber(form.EarNumber).
		SetName(form.Name).
		SetPregnancyTestMethodId(form.PregnancyTestMethodId).
		SetPregnancyTestMethodName(form.PregnancyTestMethodName).
		SetPregnancyTestResultId(form.PregnancyTestResultId).
		SetPregnancyTestResultName(form.PregnancyTestResultName).
		SetPregnancyTestTypeId(form.PregnancyTestTypeId).
		SetPregnancyTestTypeName(form.PregnancyTestTypeName).
		SetRemarks(form.Remarks).
		SetCattleId(form.CattleId).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTenantId(form.TenantId).SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetShedName(form.ShedName).
		SetTestAt(form.TestAt).
		SetTimes(form.Times).
		SetUserName(form.UserName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(pregnancytest))
}
