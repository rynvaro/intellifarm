package rations

import (
	"cattleai/ent/predicate"
	"cattleai/ent/ration"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Ration {
	wheres := []predicate.Ration{ration.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, ration.NameContains(listParams.Q))
	}
	return ration.And(wheres...)
}
