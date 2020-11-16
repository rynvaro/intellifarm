package breedings

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/breedings", BreedingAddHandler)
	g.GET("/breedings", BreedingListHandler)
	g.DELETE("/breedings/:id", BreedingDeleteHandler)
	g.PUT("/breedings/:id", BreedingUpdateHandler)
}
