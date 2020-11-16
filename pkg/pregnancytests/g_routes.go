package pregnancytests

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/pregnancytests", PregnancyTestAddHandler)
	g.GET("/pregnancytests", PregnancyTestListHandler)
	g.DELETE("/pregnancytests/:id", PregnancyTestDeleteHandler)
	g.PUT("/pregnancytests/:id", PregnancyTestUpdateHandler)
}
