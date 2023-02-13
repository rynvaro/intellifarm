package acharts

import (
	"cattleai/db"
	"cattleai/resp"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type IndexData struct {
	TenantCount int `json:"tenantCount"`
	FarmCount   int `json:"farmCount"`
	ShedCount   int `json:"shedCount"`
	UserCount   int `json:"userCount"`
}

func IndexDataHandler(c *gin.Context) {
	tenantCount, err := db.Client.Tenant.Query().Count(c)
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	farmCount, err := db.Client.Farm.Query().Count(c)
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	shedCount, err := db.Client.Shed.Query().Count(c)
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	userCount, err := db.Client.User.Query().Count(c)
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	data := IndexData{
		TenantCount: tenantCount,
		FarmCount:   farmCount,
		ShedCount:   shedCount,
		UserCount:   userCount,
	}

	c.JSON(http.StatusOK, resp.Success(data))
}
