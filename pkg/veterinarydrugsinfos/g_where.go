package veterinarydrugsinfos

import (
	"cattleai/ent/predicate"
	"cattleai/ent/veterinarydrugsinfo"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.VeterinaryDrugsInfo {
	wheres := []predicate.VeterinaryDrugsInfo{veterinarydrugsinfo.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, veterinarydrugsinfo.NameContains(listParams.Q))
	}
	wheres = append(wheres, veterinarydrugsinfo.TenantId(listParams.TenantId))
	return veterinarydrugsinfo.And(wheres...)
}
