package calves

import (
	"cattleai/ent/calve"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Calve {
	wheres := []predicate.Calve{calve.CattleId(listParams.CattleId)}
	if listParams.Q != "" {
		wheres = append(wheres, calve.NameContains(listParams.Q))
	}
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, calve.CalveAtGTE(listParams.TimeRange[0]))
		wheres = append(wheres, calve.CalveAtLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, calve.TenantId(listParams.TenantId))
	return calve.And(wheres...)
}
