package estruses

import (
	"cattleai/ent/estrus"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Estrus {
	wheres := []predicate.Estrus{estrus.CattleId(listParams.CattleId)}
	if listParams.Q != "" {
		wheres = append(wheres, estrus.NameContains(listParams.Q))
	}
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, estrus.EstrusAtGTE(listParams.TimeRange[0]))
		wheres = append(wheres, estrus.EstrusAtLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, estrus.TenantId(listParams.TenantId))
	return estrus.And(wheres...)
}
