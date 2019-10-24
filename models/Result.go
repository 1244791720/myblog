package models

import "myblog/commonconst"

type Result struct {
	Code int
	Msg  string
	Data interface{}
}

func (*Result) Success(data interface{}) *Result {
	result := new(Result)
	result.Code = commonconst.SUCCESS_CODE
	result.Msg = "请求成功"
	result.Data = data
	return result
}

func (*Result) Error() *Result {
	result := new(Result)
	result.Code = commonconst.ERROR_CODE
	result.Msg = "服务器异常"
	result.Data = nil
	return result
}

func (*Result) UnknownParam() *Result {
	result := new(Result)
	result.Code = commonconst.UNKNOW_PARAM
	result.Msg = "参数错误"
	result.Data = nil
	return result
}

func (*Result) Build(code int, msg string, data interface{}) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
