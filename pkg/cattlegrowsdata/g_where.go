package cattlegrowsdata

import (
	"cattleai/ent/cattlegrowsdata"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.CattleGrowsData {
	wheres := []predicate.CattleGrowsData{cattlegrowsdata.CattleId(listParams.CattleId)}
	if listParams.Q != "" {
		wheres = append(wheres, cattlegrowsdata.MeasuredByContains(listParams.Q))
	}
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, cattlegrowsdata.MeasuredAtGTE(listParams.TimeRange[0]))
		wheres = append(wheres, cattlegrowsdata.MeasuredAtLTE(listParams.TimeRange[1]))
	}
	wheres = append(wheres, cattlegrowsdata.TenantId(listParams.TenantId))
	return cattlegrowsdata.And(wheres...)
}
