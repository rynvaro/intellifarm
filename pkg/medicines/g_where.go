package medicines

import (
	"cattleai/ent/medicine"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Medicine {
	wheres := []predicate.Medicine{medicine.CattleId(listParams.CattleId)}
	if listParams.Epid > 0 {
		wheres = append(wheres, medicine.EpidEQ(listParams.Epid))
	}
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, medicine.DateStartGTE(listParams.TimeRange[0]))
		wheres = append(wheres, medicine.DateStartLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, medicine.TenantId(listParams.TenantId))
	return medicine.And(wheres...)
}
