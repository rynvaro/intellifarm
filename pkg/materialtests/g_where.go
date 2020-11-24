package materialtests

import (
	"cattleai/ent/materialtest"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.MaterialTest {
	wheres := []predicate.MaterialTest{materialtest.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, materialtest.NameContains(listParams.Q))
	}
	return materialtest.And(wheres...)
}
