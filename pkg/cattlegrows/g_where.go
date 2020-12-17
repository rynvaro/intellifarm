package cattlegrows

import (
	"cattleai/ent/cattlegrow"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.CattleGrow {
	wheres := []predicate.CattleGrow{cattlegrow.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, cattlegrow.NameContains(listParams.Q))
	}
	wheres = append(wheres, cattlegrow.TenantId(listParams.TenantId))
	return cattlegrow.And(wheres...)
}
