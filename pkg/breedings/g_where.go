package breedings

import (
	"cattleai/ent/breeding"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Breeding {
	wheres := []predicate.Breeding{breeding.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, breeding.NameContains(listParams.Q))
	}
	wheres = append(wheres, breeding.TenantId(listParams.TenantId))
	return breeding.And(wheres...)
}
