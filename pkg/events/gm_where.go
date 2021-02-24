package events

import (
	"cattleai/ent/event"
	"cattleai/ent/predicate"
)

func Where(listParams *EventSearchParams) predicate.Event {
	wheres := []predicate.Event{event.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, event.EarNumberEQ(listParams.Q))
	}
	if listParams.Type != "" {
		wheres = append(wheres, event.EventTypeEQ(listParams.Type))
	}
	if listParams.TimeStart > 0 {
		wheres = append(wheres, event.CreatedAtGTE(listParams.TimeStart*1000000))
	}
	if listParams.TimeEnd > 0 {
		wheres = append(wheres, event.CreatedAtLTE(listParams.TimeEnd*1000000))
	}
	wheres = append(wheres, event.TenantId(listParams.TenantId))
	return event.And(wheres...)
}
