package materials

import (
	"cattleai/ent/material"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Material {
	wheres := []predicate.Material{material.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, material.NameContains(listParams.Q))
	}
	wheres = append(wheres, material.TenantId(listParams.TenantId))
	return material.And(wheres...)
}
