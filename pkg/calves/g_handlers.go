package calves

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/calve"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CalveAddHandler(c *gin.Context) {
	form := &ent.Calve{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	calve, err := db.Client.Calve.Create().
		SetBabyBreedId(form.BabyBreedId).
		SetBabyBreedName(form.BabyBreedName).
		SetBabyEarNumber(form.BabyEarNumber).
		SetBabyGender(form.BabyGender).
		SetBabyHairColorId(form.BabyHairColorId).
		SetBabyHairColorName(form.BabyHairColorName).
		SetBabyShedId(form.BabyShedId).
		SetBabyShedName(form.BabyShedName).
		SetBabyStatus(form.BabyStatus).
		SetBabyWeight(form.BabyWeight).
		SetCalveAt(form.CalveAt).
		SetCalveCate(form.CalveCate).
		SetCalveCountId(form.CalveCountId).
		SetCalveCountName(form.CalveCountName).
		SetCalveTypeId(form.CalveTypeId).
		SetCalveTypeName(form.CalveTypeName).
		SetComplexity(form.Complexity).
		SetEarNumber(form.EarNumber).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetPregnantAt(form.PregnantAt).
		SetTimes(form.Times).
		SetUserName(form.UserName).
		SetTenantId(c.MustGet("tenantId").(int64)).
		SetTenantName(c.MustGet("tenantName").(string)).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if _, err := db.Client.Event.Create().SetCreatedAt(time.Now().UnixNano()).
		SetDeleted(0).SetEarNumber(form.EarNumber).SetEventName("产犊").
		SetEventType("繁殖事件").SetTenantId(c.MustGet("tenantId").(int64)).
		SetTenantName(c.MustGet("tenantName").(string)).Save(c.Request.Context()); err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, resp.Success(calve))
}

func CalveListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	listParams.TenantId = c.MustGet("tenantId").(int64)
	where := Where(listParams)
	totalCount, err := db.Client.Calve.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	calves, err := db.Client.Calve.Query().Where(where).Order(ent.Desc(calve.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   calves,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func CalveDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	calve, err := db.Client.Calve.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(calve))
}

func CalveUpdateHandler(c *gin.Context) {
	form := &ent.Calve{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	calve, err := db.Client.Calve.UpdateOneID(form.ID).
		SetBabyBreedId(form.BabyBreedId).
		SetBabyBreedName(form.BabyBreedName).
		SetBabyEarNumber(form.BabyEarNumber).
		SetBabyGender(form.BabyGender).
		SetBabyHairColorId(form.BabyHairColorId).
		SetBabyHairColorName(form.BabyHairColorName).
		SetBabyShedId(form.BabyShedId).
		SetBabyShedName(form.BabyShedName).
		SetBabyStatus(form.BabyStatus).
		SetBabyWeight(form.BabyWeight).
		SetCalveAt(form.CalveAt).
		SetCalveCate(form.CalveCate).
		SetCalveCountId(form.CalveCountId).
		SetCalveCountName(form.CalveCountName).
		SetCalveTypeId(form.CalveTypeId).
		SetCalveTypeName(form.CalveTypeName).
		SetComplexity(form.Complexity).
		SetEarNumber(form.EarNumber).
		SetName(form.Name).
		SetRemarks(form.Remarks).
		SetPregnantAt(form.PregnantAt).
		SetTimes(form.Times).
		SetUserName(form.UserName).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(calve))
}
