package cattleins

import (
	"cattleai/ent/cattlein"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.CattleIn {
	wheres := []predicate.CattleIn{cattlein.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, cattlein.NameContains(listParams.Q))
	}
	return cattlein.And(wheres...)
}
