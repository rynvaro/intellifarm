package tenants

import (
	"cattleai/db"
	"cattleai/ent/farm"
	"cattleai/pkg/params"
	"cattleai/resp"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func AllTenants(c *gin.Context) {
	tenants, err := db.Client.Tenant.Query().All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(tenants))
}

func TenantFarmsHandler(c *gin.Context) {
	id := &params.Id{}
	if err := c.BindUri(id); err != nil {
		log.Error().Msg(err.Error())
		return
	}
	farms, err := db.Client.Farm.Query().Where(farm.TenantId(id.Id)).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(farms))
}
