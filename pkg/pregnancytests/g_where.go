package pregnancytests

import (
	"cattleai/ent/predicate"
	"cattleai/ent/pregnancytest"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.PregnancyTest {
	wheres := []predicate.PregnancyTest{pregnancytest.CattleId(listParams.CattleId)}
	if listParams.Q != "" {
		wheres = append(wheres, pregnancytest.NameContains(listParams.Q))
	}
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, pregnancytest.TestAtGTE(listParams.TimeRange[0]))
		wheres = append(wheres, pregnancytest.TestAtLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, pregnancytest.TenantId(listParams.TenantId))
	return pregnancytest.And(wheres...)
}
