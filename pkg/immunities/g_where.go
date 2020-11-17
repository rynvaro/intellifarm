package immunities

import (
	"cattleai/ent/immunity"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Immunity {
	wheres := []predicate.Immunity{immunity.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, immunity.NameContains(listParams.Q))
	}
	return immunity.And(wheres...)
}
