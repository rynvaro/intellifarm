package sheds

import (
	"cattleai/ent/predicate"
	"cattleai/ent/shed"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Shed {
	wheres := []predicate.Shed{}
	if listParams.Q != "" {
		wheres = append(wheres, shed.NameContains(listParams.Q))
	}
	switch listParams.Level {
	case 1:
		if listParams.TenantId > 0 {
			wheres = append(wheres, shed.TenantId(listParams.TenantId))
		}
		if listParams.FarmId > 0 {
			wheres = append(wheres, shed.FarmId(listParams.FarmId))
		}
	case 2:
		wheres = append(wheres, shed.TenantId(listParams.TenantId))
		if listParams.FarmId > 0 {
			wheres = append(wheres, shed.FarmId(listParams.FarmId))
		}
	case 3:
		wheres = append(wheres, shed.FarmId(listParams.FarmId))
	}
	return shed.And(wheres...)
}
