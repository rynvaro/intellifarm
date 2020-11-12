package birthsurroundings

import (
	"cattleai/ent/birthsurrounding"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.BirthSurrounding {
	wheres := []predicate.BirthSurrounding{birthsurrounding.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, birthsurrounding.NameContains(listParams.Q))
	}
	if len(listParams.TimeRange) == 2 {
		wheres = append(wheres, birthsurrounding.RecordTimeGTE(listParams.TimeRange[0]))
		wheres = append(wheres, birthsurrounding.RecordTimeLTE(listParams.TimeRange[1]))
	}
	if listParams.UserId > 0 {
		wheres = append(wheres, birthsurrounding.UserIdEQ(listParams.UserId))
	}
	return birthsurrounding.And(wheres...)
}
