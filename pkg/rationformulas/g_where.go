package rationformulas

import (
	"cattleai/ent/predicate"
	"cattleai/ent/rationformula"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.RationFormula {
	wheres := []predicate.RationFormula{}
	if listParams.Q != "" {
		wheres = append(wheres, rationformula.NameContains(listParams.Q))
	}
	switch listParams.Level {
	case 1:
		if listParams.TenantId > 0 {
			wheres = append(wheres, rationformula.TenantId(listParams.TenantId))
		}
		if listParams.FarmId > 0 {
			wheres = append(wheres, rationformula.FarmId(listParams.FarmId))
		}
	case 2:
		wheres = append(wheres, rationformula.TenantId(listParams.TenantId))
		if listParams.FarmId > 0 {
			wheres = append(wheres, rationformula.FarmId(listParams.FarmId))
		}
	case 3:
		wheres = append(wheres, rationformula.FarmId(listParams.FarmId))
	}
	return rationformula.And(wheres...)
}
