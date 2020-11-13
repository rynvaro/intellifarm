package cattlegrowsrates

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/cattlegrowsrates", CattleGrowsRateAddHandler)
	g.GET("/cattlegrowsrates", CattleGrowsRateListHandler)
	g.DELETE("/cattlegrowsrates/:id", CattleGrowsRateDeleteHandler)
	g.PUT("/cattlegrowsrates/:id", CattleGrowsRateUpdateHandler)
}
