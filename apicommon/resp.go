package apicommon

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessResponse(data interface{}) *Resp {
	return Response(200, "OK", data)
}

func DefaultResponse() *Resp {
	return Response(200, "OK", nil)
}

func Response(code int, msg string, data interface{}) *Resp {
	return &Resp{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
