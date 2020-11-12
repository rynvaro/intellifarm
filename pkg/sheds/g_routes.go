package sheds

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/sheds", ShedAddHandler)
	g.GET("/sheds", ShedListHandler)
	g.DELETE("/sheds/:id", ShedDeleteHandler)
	g.PUT("/sheds/:id", ShedUpdateHandler)
}
