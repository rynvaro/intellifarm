package diseaseinfos

import (
	"cattleai/ent/diseaseinfo"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.DiseaseInfo {
	wheres := []predicate.DiseaseInfo{diseaseinfo.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, diseaseinfo.NameContains(listParams.Q))
	}
	wheres = append(wheres, diseaseinfo.TenantId(listParams.TenantId))
	return diseaseinfo.And(wheres...)
}
