package warehousesettings

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/warehousesettings", WarehouseSettingAddHandler)
	g.GET("/warehousesettings", WarehouseSettingListHandler)
	g.DELETE("/warehousesettings/:id", WarehouseSettingDeleteHandler)
	g.PUT("/warehousesettings/:id", WarehouseSettingUpdateHandler)
}
