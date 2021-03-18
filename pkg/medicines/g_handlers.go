package medicines

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/medicine"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func MedicineAddHandler(c *gin.Context) {
	form := &ent.Medicine{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	medicine, err := db.Client.Medicine.Create().
		SetDateEnd(form.DateEnd).
		SetDateStart(form.DateStart).
		SetDose(form.Dose).
		SetEarNumber(form.EarNumber).
		SetEpid(form.Epid).
		SetMedicineName(form.MedicineName).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetRemarks(form.Remarks).
		SetUnit(form.Unit).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if _, err := db.Client.Event.Create().SetCreatedAt(time.Now().UnixNano()).
		SetDeleted(0).SetEarNumber(form.EarNumber).SetEventName("用药").
		SetEventType("兽医事件").SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).Save(c.Request.Context()); err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp.Success(medicine))
}

func MedicineListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	where := Where(listParams)
	totalCount, err := db.Client.Medicine.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	medicines, err := db.Client.Medicine.Query().Where(where).Order(ent.Desc(medicine.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   medicines,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func MedicineDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.Medicine.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func MedicineUpdateHandler(c *gin.Context) {
	form := &ent.Medicine{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	medicine, err := db.Client.Medicine.UpdateOneID(form.ID).
		SetDateEnd(form.DateEnd).
		SetDateStart(form.DateStart).
		SetDose(form.Dose).
		SetEarNumber(form.EarNumber).
		SetEpid(form.Epid).
		SetMedicineName(form.MedicineName).
		SetRemarks(form.Remarks).
		SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).
		SetUnit(form.Unit).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(medicine))
}
