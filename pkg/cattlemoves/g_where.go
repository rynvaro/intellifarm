package cattlemoves

import (
	"cattleai/ent/cattlemove"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.CattleMove {
	wheres := []predicate.CattleMove{cattlemove.CattleId(listParams.CattleId)}
	// if listParams.Q != "" {
	// 	wheres = append(wheres, cattlemove.NameContains(listParams.Q))
	// }
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, cattlemove.DateGTE(listParams.TimeRange[0]))
		wheres = append(wheres, cattlemove.DateLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, cattlemove.TenantId(listParams.TenantId))
	return cattlemove.And(wheres...)
}
