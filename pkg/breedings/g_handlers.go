package breedings

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/breeding"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func BreedingAddHandler(c *gin.Context) {
	form := &ent.Breeding{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	breeding, err := db.Client.Breeding.Create().
		SetBreedingAt(form.BreedingAt).
		SetBreedingTypeId(form.BreedingTypeId).
		SetBreedingTypeName(form.BreedingTypeName).
		SetBullId(form.BullId).
		SetCount(form.Count).
		SetEarNumber(form.EarNumber).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetSemenFrozenTypeId(form.SemenFrozenTypeId).
		SetSemenFrozenTypeName(form.SemenFrozenTypeName).
		SetShedName(form.ShedName).
		SetTimes(form.Times).
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

	if _, err := db.Client.Event.Create().SetCreatedAt(time.Now().UnixNano()).
		SetDeleted(0).SetEarNumber(form.EarNumber).SetEventName("配种").
		SetEventType("繁殖事件").SetTenantId(form.TenantId).
		SetTenantName(form.TenantName).Save(c.Request.Context()); err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp.Success(breeding))
}

func BreedingListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.Level = c.MustGet("level").(int)
	where := Where(listParams)
	totalCount, err := db.Client.Breeding.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	breedings, err := db.Client.Breeding.Query().Where(where).Order(ent.Desc(breeding.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   breedings,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func BreedingDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	err := db.Client.Breeding.DeleteOneID(id.Id).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
}

func BreedingUpdateHandler(c *gin.Context) {
	form := &ent.Breeding{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	breeding, err := db.Client.Breeding.UpdateOneID(form.ID).
		SetBreedingAt(form.BreedingAt).
		SetBreedingTypeId(form.BreedingTypeId).
		SetBreedingTypeName(form.BreedingTypeName).
		SetBullId(form.BullId).
		SetCount(form.Count).
		SetEarNumber(form.EarNumber).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetSemenFrozenTypeId(form.SemenFrozenTypeId).
		SetSemenFrozenTypeName(form.SemenFrozenTypeName).
		SetShedName(form.ShedName).
		SetTimes(form.Times).
		SetUserName(form.UserName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(breeding))
}
