package materials

import (
	"cattleai/ent/material"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Material {
	wheres := []predicate.Material{}
	if listParams.Q != "" {
		wheres = append(wheres, material.NameContains(listParams.Q))
	}
	switch listParams.Level {
	case 1:
		if listParams.TenantId > 0 {
			wheres = append(wheres, material.TenantId(listParams.TenantId))
		}
		if listParams.FarmId > 0 {
			wheres = append(wheres, material.FarmId(listParams.FarmId))
		}
	case 2:
		wheres = append(wheres, material.TenantId(listParams.TenantId))
		if listParams.FarmId > 0 {
			wheres = append(wheres, material.FarmId(listParams.FarmId))
		}
	case 3:
		wheres = append(wheres, material.FarmId(listParams.FarmId))
	}
	return material.And(wheres...)
}
