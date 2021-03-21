package cattles

import (
	"cattleai/ent/cattle"
	"cattleai/ent/predicate"
)

func Where(listParams *CattleSearchParams) predicate.Cattle {
	wheres := []predicate.Cattle{}
	if listParams.Q != "" {
		wheres = append(wheres, cattle.EarNumberContains(listParams.Q))
	}
	if listParams.ShedId > 0 {
		wheres = append(wheres, cattle.ShedId(listParams.ShedId))
	}
	if listParams.CattleCateId > 0 {
		wheres = append(wheres, cattle.CateId(listParams.CattleCateId))
	}
	switch listParams.Level {
	case 1:
		if listParams.TenantId > 0 {
			wheres = append(wheres, cattle.TenantId(listParams.TenantId))
		}
		if listParams.FarmId > 0 {
			wheres = append(wheres, cattle.FarmId(listParams.FarmId))
		}
	case 2:
		wheres = append(wheres, cattle.TenantId(listParams.TenantId))
		if listParams.FarmId > 0 {
			wheres = append(wheres, cattle.FarmId(listParams.FarmId))
		}
	case 3:
		wheres = append(wheres, cattle.FarmId(listParams.FarmId))
	}
	return cattle.And(wheres...)
}
