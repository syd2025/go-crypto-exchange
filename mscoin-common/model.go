package mscoin_common

type BizCode int

const SuccessCode BizCode = 0

type Result struct {
	Code    BizCode `json:"code"`
	Data    any     `json:"message"`
	Message string  `json:"message"`
}

func NewResult() *Result {
	return &Result{}
}

func (r *Result) Fail(code BizCode, msg string) {
	r.Code = code
	r.Message = msg
}

func (r *Result) Success(data any) {
	r.Code = SuccessCode
	r.Message = "success"
	r.Data = data
}

func (r *Result) Deal(data any, err error) *Result {
	if err != nil {
		r.Fail(-999, err.Error())
		return r
	}
	r.Success(data)
	return r
}
