package router

import (
	"cattleai/pkg/birthsurroundings"
	"cattleai/pkg/breathrates"
	"cattleai/pkg/categories"
	"cattleai/pkg/duties"
	"cattleai/pkg/farms"
	"cattleai/pkg/hairstates"
	"cattleai/pkg/positions"
	"cattleai/pkg/shedcates"
	"cattleai/pkg/sheds"
	"cattleai/pkg/shedtypes"
	"cattleai/pkg/users"
	"cattleai/pkg/winddirections"

	"github.com/gin-gonic/gin"
)

func Register(g *gin.RouterGroup) {

	// users routes
	users.RegisterRoutes(g)
	// farms routes
	farms.RegisterRoutes(g)
	// categories routes
	categories.RegisterRoutes(g)
	// sheds routes
	sheds.RegisterRoutes(g)
	// shed cates
	shedcates.RegisterRoutes(g)
	// shed types
	shedtypes.RegisterRoutes(g)
	// breath rate
	breathrates.RegisterRoutes(g)
	// hair state
	hairstates.RegisterRoutes(g)
	// wind directions
	winddirections.RegisterRoutes(g)
	// birth surroundings
	birthsurroundings.RegisterRoutes(g)
	// positions
	positions.RegisterRoutes(g)
	// duties
	duties.RegisterRoutes(g)
}
