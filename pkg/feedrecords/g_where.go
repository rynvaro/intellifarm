package feedrecords

import (
	"cattleai/ent/feedrecord"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.FeedRecord {
	wheres := []predicate.FeedRecord{feedrecord.Deleted(0)}
	if listParams.Q != "" {
		wheres = append(wheres, feedrecord.NameContains(listParams.Q))
	}
	return feedrecord.And(wheres...)
}
