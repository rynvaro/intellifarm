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

func SaveOperation() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.String() == "/api/v1/user/login" {
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
