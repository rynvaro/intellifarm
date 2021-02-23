package positions

import (
	"cattleai/db"
	"cattleai/resp"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Positions(c *gin.Context) {
	positions, err := db.Client.Position.Query().All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(positions))
}
