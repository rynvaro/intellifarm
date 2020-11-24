package materialtests

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/materialtests", MaterialTestAddHandler)
	g.GET("/materialtests", MaterialTestListHandler)
	g.DELETE("/materialtests/:id", MaterialTestDeleteHandler)
	g.PUT("/materialtests/:id", MaterialTestUpdateHandler)
}
