package rationprocess

import (
	"cattleai/ent/predicate"
	"cattleai/ent/rationprocess"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.RationProcess {
	wheres := []predicate.RationProcess{rationprocess.RationId(listParams.Id)}
	if listParams.Q != "" {
		wheres = append(wheres, rationprocess.NameContains(listParams.Q))
	}
	wheres = append(wheres, rationprocess.TenantId(listParams.TenantId))
	return rationprocess.And(wheres...)
}
