package diseaseinfos

import "github.com/gin-gonic/gin"

func RegisterRoutes(g *gin.RouterGroup) {
	g.POST("/diseaseinfos", DiseaseInfoAddHandler)
	g.GET("/diseaseinfos", DiseaseInfoListHandler)
	g.DELETE("/diseaseinfos/:id", DiseaseInfoDeleteHandler)
	g.PUT("/diseaseinfos/:id", DiseaseInfoUpdateHandler)
}
