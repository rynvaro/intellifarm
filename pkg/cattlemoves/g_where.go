package cattlemoves

import (
	"cattleai/ent/cattlemove"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.CattleMove {
	wheres := []predicate.CattleMove{cattlemove.Deleted(0)}
	// if listParams.Q != "" {
	// 	wheres = append(wheres, cattlemove.NameContains(listParams.Q))
	// }
	wheres = append(wheres, cattlemove.TenantId(listParams.TenantId))
	return cattlemove.And(wheres...)
}
