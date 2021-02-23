package cattlehaircolors

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.GET("/cattlehaircolors", CattleHairColors)
}
