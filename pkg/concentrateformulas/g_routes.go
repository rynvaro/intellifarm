package concentrateformulas

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/concentrateformulas", ConcentrateFormulaAddHandler)
	g.GET("/concentrateformulas", ConcentrateFormulaListHandler)
	g.DELETE("/concentrateformulas/:id", ConcentrateFormulaDeleteHandler)
	g.PUT("/concentrateformulas/:id", ConcentrateFormulaUpdateHandler)
}
