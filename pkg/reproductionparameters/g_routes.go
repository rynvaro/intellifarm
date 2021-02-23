package reproductionparameters

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/reproductionparameters", ReproductionParametersAddHandler)
	g.GET("/reproductionparameters", ReproductionParametersListHandler)
	g.DELETE("/reproductionparameters/:id", ReproductionParametersDeleteHandler)
	g.PUT("/reproductionparameters/:id", ReproductionParametersUpdateHandler)
}
