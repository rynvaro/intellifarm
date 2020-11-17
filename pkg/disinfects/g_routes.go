package disinfects

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/disinfects", DisinfectAddHandler)
	g.GET("/disinfects", DisinfectListHandler)
	g.DELETE("/disinfects/:id", DisinfectDeleteHandler)
	g.PUT("/disinfects/:id", DisinfectUpdateHandler)
}
