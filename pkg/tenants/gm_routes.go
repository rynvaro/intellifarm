package tenants

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/tenants", TenantAddHandler)
	g.GET("/tenants", TenantListHandler)
	g.DELETE("/tenants/:id", TenantDeleteHandler)
	g.PUT("/tenants/:id", TenantUpdateHandler)
	g.GET("/tenantsall", AllTenants)
	g.GET("/tenants/:id/farms", TenantFarmsHandler)
}
