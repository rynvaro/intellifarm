package breedingtypes

import (
	"cattleai/apicommon"
	"cattleai/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func BreedingTypes(c *gin.Context) {
	breedingtypes, err := db.Client.BreedingType.Query().All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, apicommon.SuccessResponse(breedingtypes))
}
