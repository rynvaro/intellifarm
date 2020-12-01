package apis

import (
	"cattleai/ent/api"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.API {
	wheres := []predicate.API{api.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, api.NameContains(listParams.Q))
	}
	wheres = append(wheres, api.LevelEQ(listParams.Level))
	wheres = append(wheres, api.ParentIdEQ(listParams.ParentId))
	return api.And(wheres...)
}
