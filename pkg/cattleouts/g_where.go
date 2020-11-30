package cattleouts

import (
	"cattleai/ent/cattleout"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.CattleOut {
	wheres := []predicate.CattleOut{cattleout.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, cattleout.NameContains(listParams.Q))
	}
	return cattleout.And(wheres...)
}
