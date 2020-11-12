package farms

import (
	"cattleai/ent/farm"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Farm {
	wheres := []predicate.Farm{farm.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, farm.NameContains(listParams.Q))
	}
	return farm.And(wheres...)
}
