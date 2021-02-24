package cattles

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/cattles", CattleAddHandler)
	g.GET("/cattles", CattleListHandler)
	g.DELETE("/cattles/:id", CattleDeleteHandler)
	g.PUT("/cattles/:id", CattleUpdateHandler)
	g.GET("/cattles/init", InitHandler)
	g.GET("/cattles/pedigree/:id", PedigreeHandler)
}
