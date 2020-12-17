package inspections

import (
	"cattleai/ent/inspection"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Inspection {
	wheres := []predicate.Inspection{inspection.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, inspection.NameContains(listParams.Q))
	}
	wheres = append(wheres, inspection.TenantId(listParams.TenantId))
	return inspection.And(wheres...)
}
