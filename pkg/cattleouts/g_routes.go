package cattleouts

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/cattleouts", CattleOutAddHandler)
	g.GET("/cattleouts", CattleOutListHandler)
	g.DELETE("/cattleouts/:id", CattleOutDeleteHandler)
	g.PUT("/cattleouts/:id", CattleOutUpdateHandler)
}
