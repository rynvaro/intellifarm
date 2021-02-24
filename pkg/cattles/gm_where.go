package cattles

import (
	"cattleai/ent/cattle"
	"cattleai/ent/predicate"
)

func Where(listParams *CattleSearchParams) predicate.Cattle {
	wheres := []predicate.Cattle{cattle.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, cattle.EarNumberContains(listParams.Q))
	}
	if listParams.ShedId > 0 {
		wheres = append(wheres, cattle.ShedId(listParams.ShedId))
	}
	if listParams.CattleCateId > 0 {
		wheres = append(wheres, cattle.CateId(listParams.CattleCateId))
	}
	wheres = append(wheres, cattle.TenantId(listParams.TenantId))
	return cattle.And(wheres...)
}
