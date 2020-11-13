package cattlegrowsdata

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/cattlegrowsdata", CattleGrowsDataAddHandler)
	g.GET("/cattlegrowsdata", CattleGrowsDataListHandler)
	g.DELETE("/cattlegrowsdata/:id", CattleGrowsDataDeleteHandler)
	g.PUT("/cattlegrowsdata/:id", CattleGrowsDataUpdateHandler)
}
