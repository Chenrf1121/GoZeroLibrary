package errorx

const defaultCode = 1001

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultCodeError(code int, msg string) error {
	return &CodeError{Code: defaultCode, Msg: msg}
}
func (e *CodeError) Error() string {
	return e.Msg
}
func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
