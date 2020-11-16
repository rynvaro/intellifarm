package calvecounts

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.GET("/calvecounts", CalveCounts)
}
