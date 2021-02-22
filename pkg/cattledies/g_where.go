package cattledies

import (
	"cattleai/ent/cattledie"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.CattleDie {
	wheres := []predicate.CattleDie{cattledie.Deleted(0)}
	// if listParams.Q != "" {
	// 	wheres = append(wheres, cattledie.NameContains(listParams.Q))
	// }
	wheres = append(wheres, cattledie.TenantId(listParams.TenantId))
	return cattledie.And(wheres...)
}
