package epidemics

import (
	"cattleai/ent/epidemic"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Epidemic {
	wheres := []predicate.Epidemic{epidemic.CattleId(listParams.CattleId)}
	if listParams.Q != "" {
		wheres = append(wheres, epidemic.EarNumberContains(listParams.Q))
	}
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, epidemic.OnsetGTE(listParams.TimeRange[0]))
		wheres = append(wheres, epidemic.OnsetLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, epidemic.TenantId(listParams.TenantId))
	return epidemic.And(wheres...)
}
