package abortionreasons

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/abortionreasons", AbortionReasonAddHandler)
	g.GET("/abortionreasons", AbortionReasonListHandler)
	g.DELETE("/abortionreasons/:id", AbortionReasonDeleteHandler)
	g.PUT("/abortionreasons/:id", AbortionReasonUpdateHandler)
}
