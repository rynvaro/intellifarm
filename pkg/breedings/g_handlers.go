package breedings

import (
	"cattleai/apicommon"
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/breeding"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
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
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, apicommon.SuccessResponse(breeding))
}

func BreedingListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
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
	c.JSON(http.StatusOK, apicommon.SuccessResponse(pageData))
}

func BreedingDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	breeding, err := db.Client.Breeding.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, apicommon.SuccessResponse(breeding))
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
	c.JSON(http.StatusOK, apicommon.SuccessResponse(breeding))
}
