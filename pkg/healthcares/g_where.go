package healthcares

import (
	"cattleai/ent/healthcare"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.HealthCare {
	wheres := []predicate.HealthCare{healthcare.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, healthcare.EarNumberContains(listParams.Q))
	}
	wheres = append(wheres, healthcare.TenantId(listParams.TenantId))
	return healthcare.And(wheres...)
}
