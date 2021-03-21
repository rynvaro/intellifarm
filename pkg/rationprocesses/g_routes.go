package rationprocess

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/rationprocesses", RationProcessAddHandler)
	g.GET("/rationprocesses", RationProcessListHandler)
	g.DELETE("/rationprocesses/:id", RationProcessDeleteHandler)
	g.PUT("/rationprocesses/:id", RationProcessUpdateHandler)
}
