package users

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/user/login", UserLogin)
	g.GET("/user/info", UserInfo)
}
