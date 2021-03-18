package middleware

import (
	"cattleai/cache"
	"cattleai/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParseTenant() gin.HandlerFunc {
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
		c.Set("level", currentUser.Level)
		c.Next()
	}
}
