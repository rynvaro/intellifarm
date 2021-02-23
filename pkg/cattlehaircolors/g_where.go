package cattlehaircolors

import (
	"cattleai/ent/cattlehaircolor"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.CattleHairColor {
	wheres := []predicate.CattleHairColor{cattlehaircolor.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, cattlehaircolor.NameContains(listParams.Q))
	}
	wheres = append(wheres, cattlehaircolor.TenantId(listParams.TenantId))
	return cattlehaircolor.And(wheres...)
}
