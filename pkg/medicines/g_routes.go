package medicines

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/medicines", MedicineAddHandler)
	g.GET("/medicines", MedicineListHandler)
	g.DELETE("/medicines/:id", MedicineDeleteHandler)
	g.PUT("/medicines/:id", MedicineUpdateHandler)
}
