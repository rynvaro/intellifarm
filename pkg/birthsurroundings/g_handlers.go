package birthsurroundings

import (
	"cattleai/apicommon"
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/birthsurrounding"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func BirthSurroundingAddHandler(c *gin.Context) {
	form := &ent.BirthSurrounding{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	birthsurrounding, err := db.Client.BirthSurrounding.Create().
		SetBreathRateId(form.BreathRateId).
		SetBreathRateName(form.BreathRateName).
		SetFarmId(form.FarmId).
		SetFarmName(form.FarmName).
		SetHairStateId(form.HairStateId).
		SetHairStateName(form.HairStateName).
		SetHumidity(form.Humidity).
		SetLocationChanges(form.LocationChanges).
		SetName(form.Name).
		SetRained(form.Rained).
		SetRecordTime(form.RecordTime).
		SetRemarks(form.Remarks).
		SetSoilDepth(form.SoilDepth).
		SetSunExposure(form.SunExposure).
		SetTemperature(form.Temperature).
		SetThIndex(form.ThIndex).
		SetUserId(form.UserId).
		SetUserName(form.UserName).
		SetWalkDistance(form.WalkDistance).
		SetWindDirection(form.WindDirection).
		SetWindDirectionId(form.WindDirectionId).
		SetWindSpeed(form.WindSpeed).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, apicommon.SuccessResponse(birthsurrounding))
}

func BirthSurroundingListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	where := Where(listParams)
	totalCount, err := db.Client.BirthSurrounding.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	birthsurroundings, err := db.Client.BirthSurrounding.Query().Where(where).Order(ent.Desc(birthsurrounding.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   birthsurroundings,
		Paging: page,
	}
	c.JSON(http.StatusOK, apicommon.SuccessResponse(pageData))
}

func BirthSurroundingDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	birthsurrounding, err := db.Client.BirthSurrounding.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, apicommon.SuccessResponse(birthsurrounding))
}

func BirthSurroundingUpdateHandler(c *gin.Context) {
	form := &ent.BirthSurrounding{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	birthsurrounding, err := db.Client.BirthSurrounding.UpdateOneID(form.ID).
		SetBreathRateId(form.BreathRateId).
		SetBreathRateName(form.BreathRateName).
		SetFarmId(form.FarmId).
		SetFarmName(form.FarmName).
		SetHairStateId(form.HairStateId).
		SetHairStateName(form.HairStateName).
		SetHumidity(form.Humidity).
		SetLocationChanges(form.LocationChanges).
		SetName(form.Name).
		SetRained(form.Rained).
		SetRecordTime(form.RecordTime).
		SetRemarks(form.Remarks).
		SetSoilDepth(form.SoilDepth).
		SetSunExposure(form.SunExposure).
		SetTemperature(form.Temperature).
		SetThIndex(form.ThIndex).
		SetUserId(form.UserId).
		SetUserName(form.UserName).
		SetWalkDistance(form.WalkDistance).
		SetWindDirection(form.WindDirection).
		SetWindDirectionId(form.WindDirectionId).
		SetWindSpeed(form.WindSpeed).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, apicommon.SuccessResponse(birthsurrounding))
}
