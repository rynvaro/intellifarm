package router

import (
	"cattleai/users"

	"github.com/gin-gonic/gin"
)

func Register(g *gin.RouterGroup) {
	// register users routes to http server
	users.RegisterRoutes(g)
}
