package cattlemoves

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/cattlemoves", CattleMoveAddHandler)
	g.GET("/cattlemoves", CattleMoveListHandler)
	g.DELETE("/cattlemoves/:id", CattleMoveDeleteHandler)
	g.PUT("/cattlemoves/:id", CattleMoveUpdateHandler)
}
