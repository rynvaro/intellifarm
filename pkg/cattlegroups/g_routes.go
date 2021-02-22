package cattlegroups

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/cattlegroups", CattleGroupAddHandler)
	g.GET("/cattlegroups", CattleGroupListHandler)
	g.DELETE("/cattlegroups/:id", CattleGroupDeleteHandler)
	g.PUT("/cattlegroups/:id", CattleGroupUpdateHandler)
}
