package abortions

import (
	"cattleai/ent/abortion"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Abortion {
	wheres := []predicate.Abortion{abortion.CattleId(listParams.CattleId)}
	if listParams.Q != "" {
		wheres = append(wheres, abortion.NameContains(listParams.Q))
	}
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, abortion.AbortionAtGTE(listParams.TimeRange[0]))
		wheres = append(wheres, abortion.AbortionAtLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, abortion.TenantId(listParams.TenantId))
	return abortion.And(wheres...)
}
