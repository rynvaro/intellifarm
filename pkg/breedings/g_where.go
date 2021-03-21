package breedings

import (
	"cattleai/ent/breeding"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Breeding {
	wheres := []predicate.Breeding{breeding.CattleId(listParams.CattleId)}
	if listParams.Q != "" {
		wheres = append(wheres, breeding.NameContains(listParams.Q))
	}
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, breeding.BreedingAtGTE(listParams.TimeRange[0]))
		wheres = append(wheres, breeding.BreedingAtLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, breeding.TenantId(listParams.TenantId))
	return breeding.And(wheres...)
}
