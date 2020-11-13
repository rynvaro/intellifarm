package cattletypes

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.GET("/cattletypes", CattleTypes)
}
