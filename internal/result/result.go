package result

import "encoding/json"

type response struct {
	Code int
	Msg  string
	Data interface{}
}

func (resp *response) ToString() string {
	str, _ := json.Marshal(*resp)
	return string(str)
}

type ResultOptions func(*response)

func ResultWithData(data interface{}) ResultOptions {
	return func(r *response) {
		r.Data = data
	}
}

func ResultWithMsg(msg string) ResultOptions {
	return func(r *response) {
		r.Msg = msg
	}
}

func Result(code int, opts ...ResultOptions) response {
	res := &response{}
	for _, v := range opts {
		v(res)
	}
	return *res
}

func SueccResult(opts ...ResultOptions) response {
	res := &response{Code: 1, Msg: "执行成功"}
	for _, v := range opts {
		v(res)
	}
	return *res
}

func FailureResult(opts ...ResultOptions) response {
	res := &response{Code: 0, Msg: "执行失败"}
	for _, v := range opts {
		v(res)
	}
	return *res
}
