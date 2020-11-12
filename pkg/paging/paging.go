package paging

type Paging struct {
	CurrentPage int `json:"currentPage" form:"currentPage" required:"true"`
	PageSize    int `json:"pageSize" form:"pageSize" required:"true"`
	TotalCount  int `json:"totalCount" form:"totalCount" required:"true"`
}

type PageData struct {
	Data   interface{} `json:"data"`
	Paging Paging      `json:"paging"`
}
