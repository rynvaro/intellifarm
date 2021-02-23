package shedtypes

import (
	"cattleai/ent/predicate"
	"cattleai/ent/shedtype"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.ShedType {
	wheres := []predicate.ShedType{shedtype.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, shedtype.NameContains(listParams.Q))
	}
	wheres = append(wheres, shedtype.TenantId(listParams.TenantId))
	return shedtype.And(wheres...)
}
