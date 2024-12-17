package project_common

type BusinessCode int

type Result struct {
	Code BusinessCode `json:"code"`
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data"`
}

func (r *Result) Success(data interface{}) *Result {
	//r.Code = model.Success
	r.Code = 200
	r.Msg = "success"
	r.Data = data
	return r
}

func (r *Result) Fail(code BusinessCode, msg string) *Result {
	r.Code = code
	r.Msg = msg
	return r
}
