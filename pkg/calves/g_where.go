package calves

import (
	"cattleai/ent/calve"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Calve {
	wheres := []predicate.Calve{calve.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, calve.NameContains(listParams.Q))
	}
	wheres = append(wheres, calve.TenantId(listParams.TenantId))
	return calve.And(wheres...)
}
