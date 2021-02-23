package frozensemeninfos

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/frozensemeninfos", FrozenSemenInfoAddHandler)
	g.GET("/frozensemeninfos", FrozenSemenInfoListHandler)
	g.DELETE("/frozensemeninfos/:id", FrozenSemenInfoDeleteHandler)
	g.PUT("/frozensemeninfos/:id", FrozenSemenInfoUpdateHandler)
}
