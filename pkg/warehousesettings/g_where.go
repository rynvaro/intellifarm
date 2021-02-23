package warehousesettings

import (
	"cattleai/ent/predicate"
	"cattleai/ent/warehousesetting"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.WarehouseSetting {
	wheres := []predicate.WarehouseSetting{warehousesetting.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, warehousesetting.NameContains(listParams.Q))
	}
	wheres = append(wheres, warehousesetting.TenantId(listParams.TenantId))
	return warehousesetting.And(wheres...)
}
