package veterinarydrugsinfos

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/veterinarydrugsinfos", VeterinaryDrugsInfoAddHandler)
	g.GET("/veterinarydrugsinfos", VeterinaryDrugsInfoListHandler)
	g.DELETE("/veterinarydrugsinfos/:id", VeterinaryDrugsInfoDeleteHandler)
	g.PUT("/veterinarydrugsinfos/:id", VeterinaryDrugsInfoUpdateHandler)
}
