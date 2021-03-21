package inspections

import (
	"cattleai/ent/inspection"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Inspection {
	wheres := []predicate.Inspection{inspection.CattleId(listParams.CattleId)}
	if listParams.Q != "" {
		wheres = append(wheres, inspection.NameContains(listParams.Q))
	}
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, inspection.DateGTE(listParams.TimeRange[0]))
		wheres = append(wheres, inspection.DateLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, inspection.TenantId(listParams.TenantId))
	return inspection.And(wheres...)
}
