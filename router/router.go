package router

import (
	"cattleai/pkg/abortionreasons"
	"cattleai/pkg/abortions"
	"cattleai/pkg/abortiontypes"
	"cattleai/pkg/apis"
	"cattleai/pkg/birthsurroundings"
	"cattleai/pkg/breathrates"
	"cattleai/pkg/breedings"
	"cattleai/pkg/breedingtypes"
	"cattleai/pkg/calvecounts"
	"cattleai/pkg/calves"
	"cattleai/pkg/calvetypes"
	"cattleai/pkg/categories"
	"cattleai/pkg/cattlebreeds"
	"cattleai/pkg/cattlecates"
	"cattleai/pkg/cattledies"
	"cattleai/pkg/cattlegenders"
	"cattleai/pkg/cattlegroups"
	"cattleai/pkg/cattlegrows"
	"cattleai/pkg/cattlegrowsdata"
	"cattleai/pkg/cattlegrowsrates"
	"cattleai/pkg/cattlehaircolors"
	"cattleai/pkg/cattleins"
	"cattleai/pkg/cattlejoinedtypes"
	"cattleai/pkg/cattlemovereasons"
	"cattleai/pkg/cattlemoves"
	"cattleai/pkg/cattleouts"
	"cattleai/pkg/cattleowners"
	"cattleai/pkg/cattles"
	"cattleai/pkg/cattletypes"
	"cattleai/pkg/concentrateformulas"
	"cattleai/pkg/concentrateprocesses"
	"cattleai/pkg/customers"
	"cattleai/pkg/diseaseinfos"
	"cattleai/pkg/disinfects"
	"cattleai/pkg/duties"
	"cattleai/pkg/epidemics"
	"cattleai/pkg/epidemictypes"
	"cattleai/pkg/estruses"
	"cattleai/pkg/estrustypes"
	"cattleai/pkg/events"
	"cattleai/pkg/farms"
	"cattleai/pkg/feedinfos"
	"cattleai/pkg/feedrecords"
	"cattleai/pkg/frozensemeninfos"
	"cattleai/pkg/hairstates"
	"cattleai/pkg/healthcares"
	"cattleai/pkg/immunities"
	"cattleai/pkg/inspections"
	"cattleai/pkg/inventoryflows"
	"cattleai/pkg/materials"
	"cattleai/pkg/materialtests"
	"cattleai/pkg/medicines"
	"cattleai/pkg/operations"
	"cattleai/pkg/positions"
	"cattleai/pkg/pregnancytestmethods"
	"cattleai/pkg/pregnancytestresults"
	"cattleai/pkg/pregnancytests"
	"cattleai/pkg/pregnancytesttypes"
	"cattleai/pkg/rations"
	"cattleai/pkg/reproductionparameters"
	"cattleai/pkg/reproductivestates"
	"cattleai/pkg/semenfrozentypes"
	"cattleai/pkg/shedcates"
	"cattleai/pkg/sheds"
	"cattleai/pkg/shedsettings"
	"cattleai/pkg/shedtypes"
	"cattleai/pkg/tenants"
	"cattleai/pkg/treatmentresults"
	"cattleai/pkg/treatmentstates"
	"cattleai/pkg/users"
	"cattleai/pkg/veterinarydrugsinfos"
	"cattleai/pkg/warehousesettings"
	"cattleai/pkg/whereabouts"
	"cattleai/pkg/winddirections"

	"github.com/gin-gonic/gin"
)

func Register(g *gin.RouterGroup) {
	// users routes
	users.RegisterRoutes(g)
	users.RegisterRoutes1(g)
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
	shedtypes.RegisterRoutes1(g)
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
	positions.RegisterRoutes1(g)
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
	cattlehaircolors.RegisterRoutes1(g)
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
	// cattle grows data
	cattlegrowsdata.RegisterRoutes(g)
	// cattle grows rate
	cattlegrowsrates.RegisterRoutes(g)
	// cattle grows
	cattlegrows.RegisterRoutes(g)

	calvetypes.RegisterRoutes(g)
	calvecounts.RegisterRoutes(g)
	calves.RegisterRoutes(g)

	estruses.RegisterRoutes(g)
	estrustypes.RegisterRoutes(g)

	semenfrozentypes.RegisterRoutes(g)
	breedings.RegisterRoutes(g)

	pregnancytests.RegisterRoutes(g)
	pregnancytesttypes.RegisterRoutes(g)
	pregnancytestmethods.RegisterRoutes(g)
	pregnancytestresults.RegisterRoutes(g)

	abortions.RegisterRoutes(g)
	abortiontypes.RegisterRoutes(g)

	epidemics.RegisterRoutes(g)
	epidemictypes.RegisterRoutes(g)
	treatmentresults.RegisterRoutes(g)
	treatmentstates.RegisterRoutes(g)
	whereabouts.RegisterRoutes(g)

	inspections.RegisterRoutes(g)

	immunities.RegisterRoutes(g)

	disinfects.RegisterRoutes(g)

	materials.RegisterRoutes(g)
	materialtests.RegisterRoutes(g)
	inventoryflows.RegisterRoutes(g)

	concentrateformulas.RegisterRoutes(g)
	concentrateprocesses.RegisterRoutes(g)
	rations.RegisterRoutes(g)
	feedrecords.RegisterRoutes(g)

	cattleins.RegisterRoutes(g)
	cattleouts.RegisterRoutes(g)
	cattlegroups.RegisterRoutes(g)
	cattlemoves.RegisterRoutes(g)
	cattledies.RegisterRoutes(g)

	apis.RegisterRoutes(g)

	tenants.RegisterRoutes(g)

	events.RegisterRoutes(g)
	operations.RegisterRoutes(g)

	medicines.RegisterRoutes(g)

	healthcares.RegisterRoutes(g)

	cattlebreeds.RegisterRoutes(g)
	cattlemovereasons.RegisterRoutes(g)
	abortionreasons.RegisterRoutes(g)

	diseaseinfos.RegisterRoutes(g)
	feedinfos.RegisterRoutes(g)
	frozensemeninfos.RegisterRoutes(g)
	reproductionparameters.RegisterRoutes(g)
	veterinarydrugsinfos.RegisterRoutes(g)

	customers.RegisterRoutes(g)
	shedsettings.RegisterRoutes(g)
	warehousesettings.RegisterRoutes(g)
}
