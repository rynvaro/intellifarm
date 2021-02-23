package shedtypes

import "github.com/gin-gonic/gin"

func RegisterRoutes1(g *gin.RouterGroup) {
	g.POST("/shedtypes", ShedTypeAddHandler)
	// g.GET("/shedtypes", ShedTypeListHandler) // 注释掉 兼容旧版本
	g.DELETE("/shedtypes/:id", ShedTypeDeleteHandler)
	g.PUT("/shedtypes/:id", ShedTypeUpdateHandler)
}
