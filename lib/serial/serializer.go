package serial

import "strconv"

type CodeType = int

type ErrorSerializer struct {
	Code   CodeType      `json:"code"`
	ErrorS string        `json:"error"`
	Params []interface{} `json:"params"`
}

func (e ErrorSerializer) Error() string {
	return e.ErrorS
}

type Response struct {
	Code CodeType `json:"code"`
}

func (e Response) Error() string {
	return "ServiceCode(" + strconv.Itoa(e.Code) + ")"
}
