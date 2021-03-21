package concentrates

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/concentrates", ConcentrateAddHandler)
	g.GET("/concentrates", ConcentrateListHandler)
	g.DELETE("/concentrates/:id", ConcentrateDeleteHandler)
	g.PUT("/concentrates/:id", ConcentrateUpdateHandler)
}
