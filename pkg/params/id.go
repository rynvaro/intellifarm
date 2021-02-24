package params

type Id struct {
	Id int64 `json:"id" form:"id" uri:"id" required:"true"`
}

type IdStr struct {
	Id string `json:"id" form:"id" uri:"id" required:"true"`
}
