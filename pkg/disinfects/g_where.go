package disinfects

import (
	"cattleai/ent/disinfect"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Disinfect {
	wheres := []predicate.Disinfect{disinfect.CattleId(listParams.CattleId)}
	if listParams.Q != "" {
		wheres = append(wheres, disinfect.NameContains(listParams.Q))
	}
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, disinfect.DateGTE(listParams.TimeRange[0]))
		wheres = append(wheres, disinfect.DateLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, disinfect.TenantId(listParams.TenantId))
	return disinfect.And(wheres...)
}
