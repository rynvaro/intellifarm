package middleware

import (
	"cattleai/cache"
	"cattleai/db"
	"cattleai/resp"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var whitelist = map[string]bool{
	"/api/v1/user/login": true,
	"/api/v1/tenantsall": true,
}

func SaveOperation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if c.Request.URL.String() == "/api/v1/user/login" {
		// 	c.Next()
		// 	return
		// }
		if exist, ok := whitelist[c.Request.URL.Path]; ok && exist {
			c.Next()
			return
		}

		token := c.Request.Header.Get("X-Token")
		currentUser, err := cache.GetUserByToken(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, resp.Response(10401, "请重新登陆", nil))
			c.Abort()
			return
		}
		c.Set("tenantId", currentUser.TenantId)
		c.Set("tenantName", currentUser.TenantName)
		if _, err := db.Client.Operation.Create().
			SetUserId(currentUser.ID).
			SetUserName(currentUser.Name).
			SetTenantId(currentUser.TenantId).
			SetTenantName(currentUser.TenantName).
			SetFarmId(currentUser.FarmId).
			SetFarmName(currentUser.FarmName).
			SetDeleted(0).
			SetMethod(c.Request.Method).
			SetIP(c.ClientIP()).
			SetAPI(c.Request.URL.String()).
			SetCreatedAt(time.Now().Unix()).
			Save(c.Request.Context()); err != nil {
			log.Error().Msg(err.Error())
		}

		c.Next()
	}
}
