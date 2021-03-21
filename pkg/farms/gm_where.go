package farms

import (
	"cattleai/ent/farm"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Farm {
	wheres := []predicate.Farm{}
	if listParams.Q != "" {
		wheres = append(wheres, farm.NameContains(listParams.Q))
	}

	switch listParams.Level {
	case 1:
		if listParams.TenantId > 0 {
			wheres = append(wheres, farm.TenantId(listParams.TenantId))
		}
	case 2:
		wheres = append(wheres, farm.TenantId(listParams.TenantId))
	case 3:
		// 只能看到自己的牧场
		wheres = append(wheres, farm.ID(listParams.FarmId))
	}
	return farm.And(wheres...)
}
