package reproductionparameters

import (
	"cattleai/ent/predicate"
	"cattleai/ent/reproductionparameters"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.ReproductionParameters {
	wheres := []predicate.ReproductionParameters{reproductionparameters.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, reproductionparameters.NameContains(listParams.Q))
	}
	wheres = append(wheres, reproductionparameters.TenantId(listParams.TenantId))
	return reproductionparameters.And(wheres...)
}
