package sheds

import (
	"cattleai/ent/predicate"
	"cattleai/ent/shed"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Shed {
	wheres := []predicate.Shed{shed.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, shed.NameContains(listParams.Q))
	}
	return shed.And(wheres...)
}
