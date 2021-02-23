package cattlemovereasons

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/cattlemovereasons", CattleMoveReasonAddHandler)
	g.GET("/cattlemovereasons", CattleMoveReasonListHandler)
	g.DELETE("/cattlemovereasons/:id", CattleMoveReasonDeleteHandler)
	g.PUT("/cattlemovereasons/:id", CattleMoveReasonUpdateHandler)
}
