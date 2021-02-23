package shedsettings

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/shedsettings", ShedSettingAddHandler)
	g.GET("/shedsettings", ShedSettingListHandler)
	g.DELETE("/shedsettings/:id", ShedSettingDeleteHandler)
	g.PUT("/shedsettings/:id", ShedSettingUpdateHandler)
}
