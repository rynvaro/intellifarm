package customers

import (
	"cattleai/ent/customer"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Customer {
	wheres := []predicate.Customer{customer.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, customer.NameContains(listParams.Q))
	}
	wheres = append(wheres, customer.TenantId(listParams.TenantId))
	return customer.And(wheres...)
}
