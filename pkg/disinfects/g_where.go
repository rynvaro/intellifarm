package disinfects

import (
	"cattleai/ent/disinfect"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Disinfect {
	wheres := []predicate.Disinfect{disinfect.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, disinfect.NameContains(listParams.Q))
	}
	return disinfect.And(wheres...)
}
