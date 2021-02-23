package cattlebreeds

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/cattlebreeds", CattleBreedAddHandler)
	g.GET("/cattlebreeds", CattleBreedListHandler)
	g.DELETE("/cattlebreeds/:id", CattleBreedDeleteHandler)
	g.PUT("/cattlebreeds/:id", CattleBreedUpdateHandler)
}
