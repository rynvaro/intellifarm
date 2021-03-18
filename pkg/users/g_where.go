package users

import (
	"cattleai/ent/predicate"
	"cattleai/ent/user"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.User {
	wheres := []predicate.User{}
	if listParams.Q != "" {
		wheres = append(wheres, user.NameContains(listParams.Q))
	}
	switch listParams.Level {
	case 1:
		if listParams.TenantId > 0 {
			wheres = append(wheres, user.TenantId(listParams.TenantId))
		}
		if listParams.FarmId > 0 {
			wheres = append(wheres, user.FarmId(listParams.FarmId))
		}
	case 2:
		wheres = append(wheres, user.TenantId(listParams.TenantId))
		if listParams.FarmId > 0 {
			wheres = append(wheres, user.FarmId(listParams.FarmId))
		}
	case 3:
		wheres = append(wheres, user.FarmId(listParams.FarmId))
	}
	return user.And(wheres...)
}
