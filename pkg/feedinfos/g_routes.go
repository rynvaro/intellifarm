package feedinfos

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/feedinfos", FeedInfoAddHandler)
	g.GET("/feedinfos", FeedInfoListHandler)
	g.DELETE("/feedinfos/:id", FeedInfoDeleteHandler)
	g.PUT("/feedinfos/:id", FeedInfoUpdateHandler)
}
