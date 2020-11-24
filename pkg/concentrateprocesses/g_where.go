package concentrateprocesses

import (
	"cattleai/ent/concentrateprocess"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.ConcentrateProcess {
	wheres := []predicate.ConcentrateProcess{concentrateprocess.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, concentrateprocess.NameContains(listParams.Q))
	}
	return concentrateprocess.And(wheres...)
}
