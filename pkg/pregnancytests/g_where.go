package pregnancytests

import (
	"cattleai/ent/predicate"
	"cattleai/ent/pregnancytest"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.PregnancyTest {
	wheres := []predicate.PregnancyTest{pregnancytest.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, pregnancytest.NameContains(listParams.Q))
	}
	wheres = append(wheres, pregnancytest.TenantId(listParams.TenantId))
	return pregnancytest.And(wheres...)
}
