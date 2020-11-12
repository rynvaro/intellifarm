package farms

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/farms", FarmAddHandler)
	g.GET("/farms", FarmListHandler)
	g.DELETE("/farms/:id", FarmDeleteHandler)
	g.PUT("/farms/:id", FarmUpdateHandler)
}
