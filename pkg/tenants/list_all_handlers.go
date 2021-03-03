package tenants

import (
	"cattleai/db"
	"cattleai/ent/tenant"
	"cattleai/resp"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func AllTenants(c *gin.Context) {
	fmt.Println("shshshshshxixxox")
	tenants, err := db.Client.Tenant.Query().Where(tenant.Deleted(0)).All(c.Request.Context())
	if err != nil {
		log.Error().Msg(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp.Success(tenants))
}
