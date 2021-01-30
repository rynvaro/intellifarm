package medicines

import (
	"cattleai/ent/medicine"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Medicine {
	wheres := []predicate.Medicine{medicine.Deleted(0)}
	wheres = append(wheres, medicine.EpidEQ(listParams.Epid))
	wheres = append(wheres, medicine.TenantId(listParams.TenantId))
	return medicine.And(wheres...)
}
