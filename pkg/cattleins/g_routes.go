package cattleins

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/cattleins", CattleInAddHandler)
	g.GET("/cattleins", CattleInListHandler)
	g.DELETE("/cattleins/:id", CattleInDeleteHandler)
	g.PUT("/cattleins/:id", CattleInUpdateHandler)
}
