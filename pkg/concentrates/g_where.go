package concentrates

import (
	"cattleai/ent/concentrate"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Concentrate {
	wheres := []predicate.Concentrate{}
	if listParams.Q != "" {
		wheres = append(wheres, concentrate.NameContains(listParams.Q))
	}
	switch listParams.Level {
	case 1:
		if listParams.TenantId > 0 {
			wheres = append(wheres, concentrate.TenantId(listParams.TenantId))
		}
		if listParams.FarmId > 0 {
			wheres = append(wheres, concentrate.FarmId(listParams.FarmId))
		}
	case 2:
		wheres = append(wheres, concentrate.TenantId(listParams.TenantId))
		if listParams.FarmId > 0 {
			wheres = append(wheres, concentrate.FarmId(listParams.FarmId))
		}
	case 3:
		wheres = append(wheres, concentrate.FarmId(listParams.FarmId))
	}
	return concentrate.And(wheres...)
}
