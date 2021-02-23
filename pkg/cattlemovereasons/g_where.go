package cattlemovereasons

import (
	"cattleai/ent/cattlemovereason"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.CattleMoveReason {
	wheres := []predicate.CattleMoveReason{cattlemovereason.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, cattlemovereason.NameContains(listParams.Q))
	}
	wheres = append(wheres, cattlemovereason.TenantId(listParams.TenantId))
	return cattlemovereason.And(wheres...)
}
