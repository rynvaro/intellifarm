package concentrateprocesses

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/concentrateprocesses", ConcentrateProcessAddHandler)
	g.GET("/concentrateprocesses", ConcentrateProcessListHandler)
	g.DELETE("/concentrateprocesses/:id", ConcentrateProcessDeleteHandler)
	g.PUT("/concentrateprocesses/:id", ConcentrateProcessUpdateHandler)
}
