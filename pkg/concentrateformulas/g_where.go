package concentrateformulas

import (
	"cattleai/ent/concentrateformula"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.ConcentrateFormula {
	wheres := []predicate.ConcentrateFormula{}
	if listParams.Q != "" {
		wheres = append(wheres, concentrateformula.NameContains(listParams.Q))
	}
	switch listParams.Level {
	case 1:
		if listParams.TenantId > 0 {
			wheres = append(wheres, concentrateformula.TenantId(listParams.TenantId))
		}
		if listParams.FarmId > 0 {
			wheres = append(wheres, concentrateformula.FarmId(listParams.FarmId))
		}
	case 2:
		wheres = append(wheres, concentrateformula.TenantId(listParams.TenantId))
		if listParams.FarmId > 0 {
			wheres = append(wheres, concentrateformula.FarmId(listParams.FarmId))
		}
	case 3:
		wheres = append(wheres, concentrateformula.FarmId(listParams.FarmId))
	}
	return concentrateformula.And(wheres...)
}
