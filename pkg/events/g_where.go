package events

import (
	"cattleai/ent/event"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.Event {
	wheres := []predicate.Event{event.Deleted(0)}
	// if listParams.Q != "" {
	// 	wheres = append(wheres, event.NameContains(listParams.Q))
	// }
	wheres = append(wheres, event.TenantId(listParams.TenantId))
	return event.And(wheres...)
}
