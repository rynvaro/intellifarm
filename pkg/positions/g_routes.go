package positions

import "github.com/gin-gonic/gin"

func RegisterRoutes1(g *gin.RouterGroup) {
	g.POST("/positions", PositionAddHandler)
	// g.GET("/positions", PositionListHandler)
	g.DELETE("/positions/:id", PositionDeleteHandler)
	g.PUT("/positions/:id", PositionUpdateHandler)
}
