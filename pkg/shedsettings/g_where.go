package shedsettings

import (
	"cattleai/ent/predicate"
	"cattleai/ent/shedsetting"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.ShedSetting {
	wheres := []predicate.ShedSetting{shedsetting.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, shedsetting.NameContains(listParams.Q))
	}
	wheres = append(wheres, shedsetting.TenantId(listParams.TenantId))
	return shedsetting.And(wheres...)
}
