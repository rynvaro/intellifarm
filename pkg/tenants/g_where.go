package tenants

import (
	"cattleai/ent/predicate"
	"cattleai/ent/tenant"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Tenant {
	wheres := []predicate.Tenant{tenant.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, tenant.NameContains(listParams.Q))
	}
	return tenant.And(wheres...)
}
