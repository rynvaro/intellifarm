package resp

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) *Resp {
	return Response(200, "OK", data)
}

func Default() *Resp {
	return Response(200, "OK", nil)
}

func Response(code int, msg string, data interface{}) *Resp {
	return &Resp{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
