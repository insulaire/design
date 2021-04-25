package result

import (
	"encoding/json"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (resp *response) ToString() string {
	str, _ := json.Marshal(*resp)
	return string(str)
}

type ResultOptions func(*response)

func WithData(data interface{}) ResultOptions {
	return func(r *response) {
		r.Data = data
	}
}

type Paging struct {
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

func WithPagingData(data interface{}, count int) ResultOptions {
	return func(r *response) {
		r.Data = Paging{Count: count, Data: data}
	}
}

func WithMsg(msg string) ResultOptions {
	return func(r *response) {
		r.Msg = msg
	}
}

func new(code int, msg string, opts ...ResultOptions) response {
	res := &response{Code: code, Msg: msg}
	for _, v := range opts {
		v(res)
	}
	return *res
}

func SueccResult(opts ...ResultOptions) response {
	return new(1, "操作成功", opts...)

}

func FailureResult(opts ...ResultOptions) response {
	return new(0, "操作失败", opts...)
}
