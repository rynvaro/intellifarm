package breedingtypes

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.GET("/breedingtypes", BreedingTypes)
}
