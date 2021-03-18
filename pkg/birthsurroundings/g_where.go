package birthsurroundings

import (
	"cattleai/ent/birthsurrounding"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.BirthSurrounding {
	wheres := []predicate.BirthSurrounding{}
	if listParams.Q != "" {
		wheres = append(wheres, birthsurrounding.UserNameContains(listParams.Q))
	}
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, birthsurrounding.RecordTimeGTE(listParams.TimeRange[0]))
		wheres = append(wheres, birthsurrounding.RecordTimeLTE(listParams.TimeRange[1]))
	}
	if listParams.UserId > 0 {
		wheres = append(wheres, birthsurrounding.UserIdEQ(listParams.UserId))
	}
	switch listParams.Level {
	case 1:
		if listParams.TenantId > 0 {
			wheres = append(wheres, birthsurrounding.TenantId(listParams.TenantId))
		}
		if listParams.FarmId > 0 {
			wheres = append(wheres, birthsurrounding.FarmId(listParams.FarmId))
		}
	case 2:
		wheres = append(wheres, birthsurrounding.TenantId(listParams.TenantId))
		if listParams.FarmId > 0 {
			wheres = append(wheres, birthsurrounding.FarmId(listParams.FarmId))
		}
	case 3:
		wheres = append(wheres, birthsurrounding.FarmId(listParams.FarmId))
	}
	return birthsurrounding.And(wheres...)
}
