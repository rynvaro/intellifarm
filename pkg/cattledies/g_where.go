package cattledies

import (
	"cattleai/ent/cattledie"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.CattleDie {
	wheres := []predicate.CattleDie{cattledie.CattleId(listParams.CattleId)}
	// if listParams.Q != "" {
	// 	wheres = append(wheres, cattledie.NameContains(listParams.Q))
	// }
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, cattledie.DateGTE(listParams.TimeRange[0]))
		wheres = append(wheres, cattledie.DateLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, cattledie.TenantId(listParams.TenantId))
	return cattledie.And(wheres...)
}
