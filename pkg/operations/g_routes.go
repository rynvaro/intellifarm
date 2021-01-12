package operations

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/operations", OperationAddHandler)
	g.GET("/operations", OperationListHandler)
	g.DELETE("/operations/:id", OperationDeleteHandler)
	g.PUT("/operations/:id", OperationUpdateHandler)
}
