package inventoryflows

import (
	"cattleai/ent/inventoryflow"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.InventoryFlow {
	wheres := []predicate.InventoryFlow{inventoryflow.MaterialId(listParams.Id)}

	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, inventoryflow.DateGTE(listParams.TimeRange[0]))
		wheres = append(wheres, inventoryflow.DateLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, inventoryflow.TenantId(listParams.TenantId))
	if listParams.Type > 0 {
		wheres = append(wheres, inventoryflow.Type(listParams.Type))
	}
	return inventoryflow.And(wheres...)
}
