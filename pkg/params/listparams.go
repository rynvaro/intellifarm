package params

import (
	"cattleai/pkg/paging"
	"fmt"
)

type ListParams struct {
	Q         string        `json:"q" form:"q"`
	TimeRange []int64       `json:"timeRange" form:"timeRange"`
	UserId    int64         `json:"userId" form:"userId"`
	Paging    paging.Paging `json:"paging" form:"paging"`
	Level     int           `json:"level" form:"level"`
	ParentId  int64         `json:"parentId" form:"parentId"`
	TenantId  int64         `json:"tenantId" form:"tenantId"`
	Epid      int64         `json:"epid" form:"epid"` // 发病ID
	FarmId    int64         `json:"farmId" form:"farmId"`
	CattleId  int64         `json:"cattleId" form:"cattleId"`
	Id        int64         `json:"id" form:"id"`
	Type      int           `json:"type" form:"type"`
}

func (p *ListParams) ToString() string {
	return fmt.Sprintf("%+v", p)
}
