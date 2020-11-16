package abortions

import (
	"cattleai/ent/abortion"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Abortion {
	wheres := []predicate.Abortion{abortion.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, abortion.NameContains(listParams.Q))
	}
	return abortion.And(wheres...)
}
