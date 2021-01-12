package farms

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/farm"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func FarmAddHandler(c *gin.Context) {
	form := &ent.Farm{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())

	farm, err := db.Client.Farm.Create().
		SetCategoryId(form.CategoryId).
		SetCategoryName(form.CategoryName).
		SetCode(form.Code).
		SetConstructionDate(form.ConstructionDate).
		SetContactAddress(form.ContactAddress).
		SetContactPhone(form.ContactPhone).
		SetContactUser(form.ContactUser).
		SetDistrictCode(form.DistrictCode).
		SetDistrictName(form.DistrictName).
		SetFeedingScale(form.FeedingScale).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetShedCount(form.ShedCount).
		SetSquare(form.Square).
		SetVarietyId(form.VarietyId).
		SetVarietyName(form.VarietyName).
		SetTenantId(c.MustGet("tenantId").(int64)).
		SetTenantName(c.MustGet("tenantName").(string)).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(farm))
}

func FarmListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.TenantId = c.MustGet("tenantId").(int64)
	where := Where(listParams)
	totalCount, err := db.Client.Farm.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	farms, err := db.Client.Farm.Query().Where(where).Order(ent.Desc(farm.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   farms,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func FarmDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	farm, err := db.Client.Farm.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(farm))
}

func FarmUpdateHandler(c *gin.Context) {
	form := &ent.Farm{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	farm, err := db.Client.Farm.UpdateOneID(form.ID).
		SetCategoryId(form.CategoryId).
		SetCategoryName(form.CategoryName).
		SetCode(form.Code).
		SetConstructionDate(form.ConstructionDate).
		SetContactAddress(form.ContactAddress).
		SetContactPhone(form.ContactPhone).
		SetContactUser(form.ContactUser).
		SetDistrictCode(form.DistrictCode).
		SetDistrictName(form.DistrictName).
		SetFeedingScale(form.FeedingScale).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetShedCount(form.ShedCount).
		SetSquare(form.Square).
		SetVarietyId(form.VarietyId).
		SetVarietyName(form.VarietyName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(farm))
}
