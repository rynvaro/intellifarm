package calves

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/calves", CalveAddHandler)
	g.GET("/calves", CalveListHandler)
	g.DELETE("/calves/:id", CalveDeleteHandler)
	g.PUT("/calves/:id", CalveUpdateHandler)
}
