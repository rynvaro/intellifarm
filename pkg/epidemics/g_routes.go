package epidemics

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/epidemics", EpidemicAddHandler)
	g.GET("/epidemics", EpidemicListHandler)
	g.DELETE("/epidemics/:id", EpidemicDeleteHandler)
	g.PUT("/epidemics/:id", EpidemicUpdateHandler)
}
