package cattlegrowsrates

import (
	"cattleai/ent/cattlegrowsrate"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.CattleGrowsRate {
	wheres := []predicate.CattleGrowsRate{cattlegrowsrate.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, cattlegrowsrate.NameContains(listParams.Q))
	}
	return cattlegrowsrate.And(wheres...)
}
