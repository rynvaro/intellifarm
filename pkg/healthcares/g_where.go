package healthcares

import (
	"cattleai/ent/healthcare"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.HealthCare {
	wheres := []predicate.HealthCare{healthcare.CattleId(listParams.CattleId)}
	if listParams.Q != "" {
		wheres = append(wheres, healthcare.EarNumberContains(listParams.Q))
	}
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, healthcare.DateGTE(listParams.TimeRange[0]))
		wheres = append(wheres, healthcare.DateLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, healthcare.TenantId(listParams.TenantId))
	return healthcare.And(wheres...)
}
