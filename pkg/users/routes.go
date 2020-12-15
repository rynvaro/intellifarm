package users

import "github.com/gin-gonic/gin"

func RegisterRoutes1(g *gin.RouterGroup) {
	g.POST("/user/login", UserLogin)
	g.GET("/user/info", UserInfo)
	g.GET("/user/logout", UserLogout)
}
