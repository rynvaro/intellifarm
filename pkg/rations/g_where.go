package rations

import (
	"cattleai/ent/predicate"
	"cattleai/ent/ration"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Ration {
	wheres := []predicate.Ration{}
	if listParams.Q != "" {
		wheres = append(wheres, ration.NameContains(listParams.Q))
	}
	switch listParams.Level {
	case 1:
		if listParams.TenantId > 0 {
			wheres = append(wheres, ration.TenantId(listParams.TenantId))
		}
		if listParams.FarmId > 0 {
			wheres = append(wheres, ration.FarmId(listParams.FarmId))
		}
	case 2:
		wheres = append(wheres, ration.TenantId(listParams.TenantId))
		if listParams.FarmId > 0 {
			wheres = append(wheres, ration.FarmId(listParams.FarmId))
		}
	case 3:
		wheres = append(wheres, ration.FarmId(listParams.FarmId))
	}
	return ration.And(wheres...)
}
