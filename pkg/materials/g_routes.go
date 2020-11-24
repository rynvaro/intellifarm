package materials

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/materials", MaterialAddHandler)
	g.GET("/materials", MaterialListHandler)
	g.DELETE("/materials/:id", MaterialDeleteHandler)
	g.PUT("/materials/:id", MaterialUpdateHandler)
}
