package shedtypes

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.GET("/shedtypes", ShedTypes)
}
