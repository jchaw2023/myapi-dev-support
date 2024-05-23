package web

// swagger:model
type ApiResult struct {
	//API请求后，业务处理成功true,异常false,同时msg为异常消息
	Success bool `json:"success"`
	//API请求后，返回的业务数据
	Data interface{} `json:"data"`
	//API请求后，返回的提示信息，如果success为false,这里为异常提示消息
	Msg string `json:"msg"`
}

func Ok(data interface{}, msg ...string) *ApiResult {
	outMsg := "操作成功!"
	if len(msg) > 0 {
		outMsg = msg[0]
	}

	r := &ApiResult{
		Success: true,
		Data:    data,
		Msg:     outMsg,
	}
	return r
}
func DefOk(msg ...string) *ApiResult {
	return Ok(struct {
	}{}, msg...)
}
func FailAny(msg string) *ApiResult {
	r := &ApiResult{
		Success: false,
		Msg:     msg,
	}
	return r
}
func Fail(msg ...string) *ApiResult {
	outMsg := "操作失败!"
	if len(msg) > 0 {
		outMsg = msg[0]
	}
	r := &ApiResult{
		Success: false,
		Msg:     outMsg,
	}
	return r
}
