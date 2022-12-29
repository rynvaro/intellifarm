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
		SetDeleted(0).SetEarNumber(form.EarNumber).SetEventTypeName("产犊").
		SetEventSubTypeName("繁殖事件").SetTenantId(form.TenantId).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTimes(form.Times).
		SetTimes(form.Times).
		SetTenantName(form.TenantName).Save(c.Request.Context()); err != nil {
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
	listParams.Level = c.MustGet("level").(int)
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
	err := db.Client.Calve.DeleteOneID(int(id.Id)).Exec(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(nil))
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
		SetCattleId(form.CattleId).
		SetShedId(form.ShedId).SetShedName(form.ShedName).
		SetTenantId(form.TenantId).SetTenantName(form.TenantName).
		SetFarmId(form.FarmId).SetFarmName(form.FarmName).
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
