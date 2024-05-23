package pagin_model

// swagger:model
type PaginParam struct {
	//分页参数页索引（1）
	Page int `json:"page"  validate:"min=1,required"  message:"起始页面大小为1"`
	//分页参数页面大小（1）
	Limit int `json:"limit"  validate:"min=1,required" message:"页面大小最小为1"`
}

func (p *PaginParam) Offset() int {
	return (p.Page - 1) * p.Limit
}

// swagger:model
type PaginData struct {
	//分页总记录数
	Count int64 `json:"count"`
	//状态码 一般为0
	Code int64 `json:"code"`
	//状态消息 正常情况为""
	Msg string `json:"msg"`
	//当前页面数据
	Data interface{} `json:"data"`
}

// NilPaginData 空数据分页列表
var NilPaginData = &PaginData{
	Data: make([]interface{}, 0),
}

// DefPaginData 默认正常的数据分页列表
func DefPaginData(count int64, data any) *PaginData {
	return &PaginData{
		Count: count,
		Code:  0,
		Data:  data,
	}
}
