package users

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/users", UserAddHandler)
	g.GET("/users", UserListHandler)
	g.DELETE("/users/:id", UserDeleteHandler)
	g.PUT("/users/:id", UserUpdateHandler)
}
