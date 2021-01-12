package epidemics

import (
	"cattleai/ent/epidemic"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Epidemic {
	wheres := []predicate.Epidemic{epidemic.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, epidemic.EarNumberContains(listParams.Q))
	}
	wheres = append(wheres, epidemic.TenantId(listParams.TenantId))
	return epidemic.And(wheres...)
}
