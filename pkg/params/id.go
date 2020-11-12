package params

type Id struct {
	Id int64 `json:"id" form:"id" uri:"id" required:"true"`
}
