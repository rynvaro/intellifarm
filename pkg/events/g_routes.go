package events

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/events", EventAddHandler)
	g.GET("/events", EventListHandler)
	g.DELETE("/events/:id", EventDeleteHandler)
	g.PUT("/events/:id", EventUpdateHandler)
}
