package feedrecords

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/feedrecords", FeedRecordAddHandler)
	g.GET("/feedrecords", FeedRecordListHandler)
	g.DELETE("/feedrecords/:id", FeedRecordDeleteHandler)
	g.PUT("/feedrecords/:id", FeedRecordUpdateHandler)
}
