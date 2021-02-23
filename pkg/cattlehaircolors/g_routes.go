package cattlehaircolors

import "github.com/gin-gonic/gin"

func RegisterRoutes1(g *gin.RouterGroup) {
	g.POST("/cattlehaircolors", CattleHairColorAddHandler)
	// g.GET("/cattlehaircolors", CattleHairColorListHandler)
	g.DELETE("/cattlehaircolors/:id", CattleHairColorDeleteHandler)
	g.PUT("/cattlehaircolors/:id", CattleHairColorUpdateHandler)
}
