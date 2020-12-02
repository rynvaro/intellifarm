package cattles

import (
	"cattleai/db"
	"cattleai/ent"
	"cattleai/ent/cattle"
	"cattleai/pkg/paging"
	"cattleai/pkg/params"
	"cattleai/resp"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CattleAddHandler(c *gin.Context) {
	form := &ent.Cattle{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattle, err := db.Client.Cattle.Create().
		SetBirthday(form.Birthday).
		SetBreed(form.Breed).
		SetBreedingAt(form.BreedingAt).
		SetBreedingTypeId(form.BreedingTypeId).
		SetBreedingTypeName(form.BreedingTypeName).
		SetBullId(form.BullId).
		SetCateId(form.CateId).
		SetCateName(form.CateName).
		SetEarNumber(form.EarNumber).
		SetElectronicEarNumber(form.ElectronicEarNumber).
		SetFarmId(form.FarmId).
		SetFarmName(form.FarmName).
		SetFather(form.Father).
		SetFrom(form.From).
		SetGenderId(form.GenderId).
		SetGenderName(form.GenderName).
		SetGrandfather(form.Grandfather).
		SetHairColorId(form.HairColorId).
		SetHairColorName(form.HairColorName).
		SetJoinedAt(form.JoinedAt).
		SetJoinedTypeId(form.JoinedTypeId).
		SetJoinedTypeName(form.JoinedTypeName).
		SetLastCalvingAt(form.LastCalvingAt).
		SetMother(form.Mother).
		SetOwnerId(form.OwnerId).
		SetOwnerName(form.OwnerName).
		SetPedometer(form.Pedometer).
		SetPregnancyCheckAt(form.PregnancyCheckAt).
		SetPregnantTimes(form.PregnantTimes).
		SetReproductiveStateId(form.ReproductiveStateId).
		SetReproductiveStateName(form.ReproductiveStateName).
		SetShedTypeId(form.ShedTypeId).
		SetShedTypeName(form.ShedTypeName).
		SetShedId(form.ShedId).
		SetShedName(form.ShedName).
		SetType(form.Type).
		SetTypeName(form.TypeName).
		SetWeight(form.Weight).
		SetCreatedAt(time.Now().Unix()).SetUpdatedAt(time.Now().Unix()).SetDeleted(0).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattle))
}

func CattleListHandler(c *gin.Context) {
	listParams := &params.ListParams{}
	if err := c.BindQuery(listParams); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	page := listParams.Paging
	where := Where(listParams)
	totalCount, err := db.Client.Cattle.Query().Where(where).Count(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	cattles, err := db.Client.Cattle.Query().Where(where).Order(ent.Desc(cattle.FieldCreatedAt)).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	page.TotalCount = totalCount
	pageData := paging.PageData{
		Data:   cattles,
		Paging: page,
	}
	c.JSON(http.StatusOK, resp.Success(pageData))
}

func CattleDeleteHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(fmt.Sprintf("%+v", id))
	cattle, err := db.Client.Cattle.UpdateOneID(id.Id).SetDeleted(1).Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattle))
}

func CattleUpdateHandler(c *gin.Context) {
	form := &ent.Cattle{}
	if err := c.Bind(form); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	log.Debug().Msg(form.String())
	cattle, err := db.Client.Cattle.UpdateOneID(form.ID).
		SetBirthday(form.Birthday).
		SetBreed(form.Breed).
		SetBreedingAt(form.BreedingAt).
		SetBreedingTypeId(form.BreedingTypeId).
		SetBreedingTypeName(form.BreedingTypeName).
		SetBullId(form.BullId).
		SetCateId(form.CateId).
		SetCateName(form.CateName).
		SetEarNumber(form.EarNumber).
		SetElectronicEarNumber(form.ElectronicEarNumber).
		SetFarmId(form.FarmId).
		SetFarmName(form.FarmName).
		SetFather(form.Father).
		SetFrom(form.From).
		SetGenderId(form.GenderId).
		SetGenderName(form.GenderName).
		SetGrandfather(form.Grandfather).
		SetHairColorId(form.HairColorId).
		SetHairColorName(form.HairColorName).
		SetJoinedAt(form.JoinedAt).
		SetJoinedTypeId(form.JoinedTypeId).
		SetJoinedTypeName(form.JoinedTypeName).
		SetLastCalvingAt(form.LastCalvingAt).
		SetMother(form.Mother).
		SetOwnerId(form.OwnerId).
		SetOwnerName(form.OwnerName).
		SetPedometer(form.Pedometer).
		SetPregnancyCheckAt(form.PregnancyCheckAt).
		SetPregnantTimes(form.PregnantTimes).
		SetReproductiveStateId(form.ReproductiveStateId).
		SetReproductiveStateName(form.ReproductiveStateName).
		SetShedTypeId(form.ShedTypeId).
		SetShedTypeName(form.ShedTypeName).
		SetShedId(form.ShedId).
		SetShedName(form.ShedName).
		SetType(form.Type).
		SetTypeName(form.TypeName).
		SetWeight(form.Weight).
		SetUpdatedAt(time.Now().Unix()).
		Save(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(cattle))
}
