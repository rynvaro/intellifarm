package cattles

import (
	"cattleai/ent/cattle"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Cattle {
	wheres := []predicate.Cattle{cattle.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, cattle.NameContains(listParams.Q))
	}
	return cattle.And(wheres...)
}
