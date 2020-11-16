package estruses

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/estruses", EstrusAddHandler)
	g.GET("/estruses", EstrusListHandler)
	g.DELETE("/estruses/:id", EstrusDeleteHandler)
	g.PUT("/estruses/:id", EstrusUpdateHandler)
}
