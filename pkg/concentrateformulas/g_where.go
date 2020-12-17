package concentrateformulas

import (
	"cattleai/ent/concentrateformula"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.ConcentrateFormula {
	wheres := []predicate.ConcentrateFormula{concentrateformula.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, concentrateformula.NameContains(listParams.Q))
	}
	wheres = append(wheres, concentrateformula.TenantId(listParams.TenantId))
	return concentrateformula.And(wheres...)
}
