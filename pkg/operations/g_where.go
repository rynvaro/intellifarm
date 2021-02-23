package operations

import (
	"cattleai/ent/operation"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Operation {
	wheres := []predicate.Operation{operation.Deleted(0)}
	// if listParams.Q != "" {
	// 	wheres = append(wheres, operation.NameContains(listParams.Q))
	// }
	wheres = append(wheres, operation.TenantId(listParams.TenantId))
	return operation.And(wheres...)
}
