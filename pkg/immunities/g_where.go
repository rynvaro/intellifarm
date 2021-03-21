package immunities

import (
	"cattleai/ent/immunity"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Immunity {
	wheres := []predicate.Immunity{immunity.CattleId(listParams.CattleId)}
	if listParams.Q != "" {
		wheres = append(wheres, immunity.NameContains(listParams.Q))
	}
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, immunity.DateGTE(listParams.TimeRange[0]))
		wheres = append(wheres, immunity.DateLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, immunity.TenantId(listParams.TenantId))
	return immunity.And(wheres...)
}
