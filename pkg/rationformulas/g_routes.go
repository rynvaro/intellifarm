package rationformulas

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/rationformulas", RationFormulaAddHandler)
	g.GET("/rationformulas", RationFormulaListHandler)
	g.DELETE("/rationformulas/:id", RationFormulaDeleteHandler)
	g.PUT("/rationformulas/:id", RationFormulaUpdateHandler)
}
