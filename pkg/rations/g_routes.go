package rations

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/rations", RationAddHandler)
	g.GET("/rations", RationListHandler)
	g.DELETE("/rations/:id", RationDeleteHandler)
	g.PUT("/rations/:id", RationUpdateHandler)
}
