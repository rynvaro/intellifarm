package inventoryflows

import (
	"cattleai/ent/inventoryflow"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.InventoryFlow {
	wheres := []predicate.InventoryFlow{inventoryflow.Deleted(0)}
	// if listParams.Q != "" {
	// 	wheres = append(wheres, inventoryflow.NameContains(listParams.Q))
	// }
	wheres = append(wheres, inventoryflow.TenantId(listParams.TenantId))
	return inventoryflow.And(wheres...)
}
