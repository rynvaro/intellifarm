package frozensemeninfos

import (
	"cattleai/ent/frozensemeninfo"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.FrozenSemenInfo {
	wheres := []predicate.FrozenSemenInfo{frozensemeninfo.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, frozensemeninfo.NameContains(listParams.Q))
	}
	wheres = append(wheres, frozensemeninfo.TenantId(listParams.TenantId))
	return frozensemeninfo.And(wheres...)
}
