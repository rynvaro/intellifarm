package hairstates

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.GET("/hairstates", HairStates)
}
