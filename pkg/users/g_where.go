package users

import (
	"cattleai/ent/predicate"
	"cattleai/ent/user"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.User {
	wheres := []predicate.User{user.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, user.NameContains(listParams.Q))
	}
	wheres = append(wheres, user.TenantId(listParams.TenantId))
	return user.And(wheres...)
}
