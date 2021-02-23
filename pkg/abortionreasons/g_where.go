package abortionreasons

import (
	"cattleai/ent/abortionreason"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.AbortionReason {
	wheres := []predicate.AbortionReason{abortionreason.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, abortionreason.NameContains(listParams.Q))
	}
	wheres = append(wheres, abortionreason.TenantId(listParams.TenantId))
	return abortionreason.And(wheres...)
}
