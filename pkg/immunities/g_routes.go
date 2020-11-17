package immunities

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/immunities", ImmunityAddHandler)
	g.GET("/immunities", ImmunityListHandler)
	g.DELETE("/immunities/:id", ImmunityDeleteHandler)
	g.PUT("/immunities/:id", ImmunityUpdateHandler)
}
