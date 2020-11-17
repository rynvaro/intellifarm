package treatmentstates

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.GET("/treatmentstates", TreatmentStates)
}
