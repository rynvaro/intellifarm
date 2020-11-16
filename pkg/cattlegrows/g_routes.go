package cattlegrows

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/cattlegrows", CattleGrowAddHandler)
	g.GET("/cattlegrows", CattleGrowListHandler)
	g.DELETE("/cattlegrows/:id", CattleGrowDeleteHandler)
	g.PUT("/cattlegrows/:id", CattleGrowUpdateHandler)
}
