package abortions

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/abortions", AbortionAddHandler)
	g.GET("/abortions", AbortionListHandler)
	g.DELETE("/abortions/:id", AbortionDeleteHandler)
	g.PUT("/abortions/:id", AbortionUpdateHandler)
}
