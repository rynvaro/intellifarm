package positions

import (
	"cattleai/ent/position"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Position {
	wheres := []predicate.Position{position.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, position.NameContains(listParams.Q))
	}
	wheres = append(wheres, position.TenantId(listParams.TenantId))
	return position.And(wheres...)
}
