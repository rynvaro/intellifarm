package cattlebreeds

import (
	"cattleai/ent/cattlebreed"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.CattleBreed {
	wheres := []predicate.CattleBreed{cattlebreed.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, cattlebreed.NameContains(listParams.Q))
	}
	wheres = append(wheres, cattlebreed.TenantId(listParams.TenantId))
	return cattlebreed.And(wheres...)
}
