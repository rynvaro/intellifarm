package inventoryflows

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/inventoryflows", InventoryFlowAddHandler)
	g.GET("/inventoryflows", InventoryFlowListHandler)
	g.DELETE("/inventoryflows/:id", InventoryFlowDeleteHandler)
	g.PUT("/inventoryflows/:id", InventoryFlowUpdateHandler)
}
