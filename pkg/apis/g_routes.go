package apis

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/apis", APIAddHandler)
	g.GET("/apis", APIListHandler)
	g.DELETE("/apis/:id", APIDeleteHandler)
	g.PUT("/apis/:id", APIUpdateHandler)
}
