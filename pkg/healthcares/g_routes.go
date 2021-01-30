package healthcares

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/healthcares", HealthCareAddHandler)
	g.GET("/healthcares", HealthCareListHandler)
	g.DELETE("/healthcares/:id", HealthCareDeleteHandler)
	g.PUT("/healthcares/:id", HealthCareUpdateHandler)
}
