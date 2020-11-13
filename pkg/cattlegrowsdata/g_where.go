package cattlegrowsdata

import (
	"cattleai/ent/cattlegrowsdata"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.CattleGrowsData {
	wheres := []predicate.CattleGrowsData{cattlegrowsdata.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, cattlegrowsdata.NameContains(listParams.Q))
	}
	return cattlegrowsdata.And(wheres...)
}
