package params

import "cattleai/pkg/paging"

type ListParams struct {
	Q         string        `json:"q" form:"q"`
	TimeRange []int64       `json:"timeRange" form:"timeRange"`
	UserId    int64         `json:"userId" form:"userId"`
	Paging    paging.Paging `json:"paging" form:"paging"`
}
