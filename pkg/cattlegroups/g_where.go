package cattlegroups

import (
	"cattleai/ent/cattlegroup"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.CattleGroup {
	wheres := []predicate.CattleGroup{cattlegroup.Deleted(0)}
	// if listParams.Q != "" {
	// 	wheres = append(wheres, cattlegroup.NameContains(listParams.Q))
	// }
	wheres = append(wheres, cattlegroup.TenantId(listParams.TenantId))
	return cattlegroup.And(wheres...)
}
