package router

import (
	"cattleai/pkg/birthsurroundings"
	"cattleai/pkg/breathrates"
	"cattleai/pkg/breedingtypes"
	"cattleai/pkg/categories"
	"cattleai/pkg/cattlecates"
	"cattleai/pkg/cattlegenders"
	"cattleai/pkg/cattlehaircolors"
	"cattleai/pkg/cattlejoinedtypes"
	"cattleai/pkg/cattleowners"
	"cattleai/pkg/cattles"
	"cattleai/pkg/cattletypes"
	"cattleai/pkg/duties"
	"cattleai/pkg/farms"
	"cattleai/pkg/hairstates"
	"cattleai/pkg/positions"
	"cattleai/pkg/reproductivestates"
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
	// breeding types
	breedingtypes.RegisterRoutes(g)
	// cattle cates
	cattlecates.RegisterRoutes(g)
	// cattle genders
	cattlegenders.RegisterRoutes(g)
	// cattle hair colors
	cattlehaircolors.RegisterRoutes(g)
	// cattle joined type
	cattlejoinedtypes.RegisterRoutes(g)
	// cattle owners
	cattleowners.RegisterRoutes(g)
	// cattle types
	cattletypes.RegisterRoutes(g)
	// reproductive states
	reproductivestates.RegisterRoutes(g)
	// cattles
	cattles.RegisterRoutes(g)
}
