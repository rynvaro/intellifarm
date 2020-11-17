package inspections

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/inspections", InspectionAddHandler)
	g.GET("/inspections", InspectionListHandler)
	g.DELETE("/inspections/:id", InspectionDeleteHandler)
	g.PUT("/inspections/:id", InspectionUpdateHandler)
}
