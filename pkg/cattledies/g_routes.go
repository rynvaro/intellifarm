package cattledies

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/cattledies", CattleDieAddHandler)
	g.GET("/cattledies", CattleDieListHandler)
	g.DELETE("/cattledies/:id", CattleDieDeleteHandler)
	g.PUT("/cattledies/:id", CattleDieUpdateHandler)
}
