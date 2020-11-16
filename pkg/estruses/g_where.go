package estruses

import (
	"cattleai/ent/estrus"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Estrus {
	wheres := []predicate.Estrus{estrus.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, estrus.NameContains(listParams.Q))
	}
	return estrus.And(wheres...)
}
