package customers

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/customers", CustomerAddHandler)
	g.GET("/customers", CustomerListHandler)
	g.DELETE("/customers/:id", CustomerDeleteHandler)
	g.PUT("/customers/:id", CustomerUpdateHandler)
}
