package feedinfos

import (
	"cattleai/ent/feedinfo"
	"cattleai/ent/predicate"
	"cattleai/pkg/params"
)

func Where(listParams *params.ListParams) predicate.FeedInfo {
	wheres := []predicate.FeedInfo{}
	if listParams.Q != "" {
		wheres = append(wheres, feedinfo.NameContains(listParams.Q))
	}
	// wheres = append(wheres, feedinfo.TenantId(listParams.TenantId))
	return feedinfo.And(wheres...)
}
