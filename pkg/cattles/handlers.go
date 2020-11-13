package cattles

import (
	"cattleai/apicommon"
	"cattleai/db"
	"cattleai/ent"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type InitData struct {
	CattleGenders      []*ent.CattleGender      `json:"cattleGenders"`
	CattleCates        []*ent.CattleCate        `json:"cattleCates"`
	CattleTypes        []*ent.CattleType        `json:"cattleTypes"`
	CattleJoinedTypes  []*ent.CattleJoinedType  `json:"cattleJoinedTypes"`
	CattleOwners       []*ent.CattleOwner       `json:"cattleOwners"`
	CattleHairColors   []*ent.CattleHairColor   `json:"cattleHairColors"`
	ReproductiveStates []*ent.ReproductiveState `json:"reproductiveStates"`
	BreedingTypes      []*ent.BreedingType      `json:"breedingTypes"`
}

func InitHandler(c *gin.Context) {
	initData := &InitData{}
	cattleGenders, err := db.Client.CattleGender.Query().All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	initData.CattleGenders = cattleGenders

	cattleCates, err := db.Client.CattleCate.Query().All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	initData.CattleCates = cattleCates

	cattleTypes, err := db.Client.CattleType.Query().All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	initData.CattleTypes = cattleTypes

	cattleJoinedTypes, err := db.Client.CattleJoinedType.Query().All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	initData.CattleJoinedTypes = cattleJoinedTypes

	cattleOwners, err := db.Client.CattleOwner.Query().All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	initData.CattleOwners = cattleOwners

	cattleHairColors, err := db.Client.CattleHairColor.Query().All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	initData.CattleHairColors = cattleHairColors

	reproductiveStates, err := db.Client.ReproductiveState.Query().All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	initData.ReproductiveStates = reproductiveStates

	breedingTypes, err := db.Client.BreedingType.Query().All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	initData.BreedingTypes = breedingTypes

	c.JSON(http.StatusOK, apicommon.SuccessResponse(initData))
}
