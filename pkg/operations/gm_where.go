package operations

import (
	"cattleai/ent/operation"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Operation {
	wheres := []predicate.Operation{}
	if listParams.Q != "" {
		wheres = append(wheres, operation.APIContains(listParams.Q))
	}
	switch listParams.Level {
	case 1:
		if listParams.TenantId > 0 {
			wheres = append(wheres, operation.TenantId(listParams.TenantId))
		}
		if listParams.FarmId > 0 {
			wheres = append(wheres, operation.FarmId(listParams.FarmId))
		}
	case 2:
		wheres = append(wheres, operation.TenantId(listParams.TenantId))
		if listParams.FarmId > 0 {
			wheres = append(wheres, operation.FarmId(listParams.FarmId))
		}
	case 3:
		wheres = append(wheres, operation.FarmId(listParams.FarmId))
	}
	return operation.And(wheres...)
}
