package birthsurroundings

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/birthsurroundings", BirthSurroundingAddHandler)
	g.GET("/birthsurroundings", BirthSurroundingListHandler)
	g.DELETE("/birthsurroundings/:id", BirthSurroundingDeleteHandler)
	g.PUT("/birthsurroundings/:id", BirthSurroundingUpdateHandler)
}
