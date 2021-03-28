package events

import (
	"cattleai/ent/event"
	"cattleai/ent/predicate"

	"github.com/rs/zerolog/log"
)

func Where(listParams *EventSearchParams) predicate.Event {
	log.Debug().Msg(listParams.ToString())

	if len(listParams.TimeRange) == 2 {
		listParams.TimeStart = listParams.TimeRange[0]
		listParams.TimeEnd = listParams.TimeRange[1]
	}

	wheres := []predicate.Event{}
	if listParams.Q != "" {
		wheres = append(wheres, event.Or(
			event.EarNumberContains(listParams.Q),
			event.EventTypeNameContains(listParams.Q),
			event.EventSubTypeNameContains(listParams.Q),
		))
	}
	if listParams.TypeName != "" {
		wheres = append(wheres, event.EventTypeNameEQ(listParams.TypeName))
	}
	if listParams.SubTypeName != "" {
		wheres = append(wheres, event.EventSubTypeNameEQ(listParams.SubTypeName))
	}
	if listParams.TimeStart > 0 {
		wheres = append(wheres, event.CreatedAtGTE(listParams.TimeStart*1000000))
	}
	if listParams.TimeEnd > 0 {
		wheres = append(wheres, event.CreatedAtLTE(listParams.TimeEnd*1000000))
	}
	if listParams.PStart > 0 {
		wheres = append(wheres, event.TimesGTE(listParams.PStart))
	}
	if listParams.PEnd > 0 {
		wheres = append(wheres, event.TimesLTE(listParams.PEnd))
	}
	switch listParams.Level {
	case 1:
		if listParams.TenantId > 0 {
			wheres = append(wheres, event.TenantId(listParams.TenantId))
		}
		if listParams.FarmId > 0 {
			wheres = append(wheres, event.FarmId(listParams.FarmId))
		}
	case 2:
		wheres = append(wheres, event.TenantId(listParams.TenantId))
		if listParams.FarmId > 0 {
			wheres = append(wheres, event.FarmId(listParams.FarmId))
		}
	case 3:
		wheres = append(wheres, event.FarmId(listParams.FarmId))
	}
	return event.And(wheres...)
}
