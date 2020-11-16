package abortiontypes

import (
	"cattleai/apicommon"
	"cattleai/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func AbortionTypes(c *gin.Context) {
	abortiontypes, err := db.Client.AbortionType.Query().All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, apicommon.SuccessResponse(abortiontypes))
}
