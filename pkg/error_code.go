package pkg

import "fmt"

type ICode interface {
	HTTPCode() int
	Message() string
	Code() int
}

type ErrorCode struct {
	status int
	code   string
	msg    string
}

func (code ErrorCode) Error() string {
	return fmt.Sprintf("GetCode: %d, msg: %s", code.code, code.msg)
}

func (code ErrorCode) Code() string {
	return code.code
}

func (code ErrorCode) Message() string {
	return code.msg
}

func (code ErrorCode) HTTPCode() int {
	return code.status
}

func Errorf(status int, code string, msg string, v ...interface{}) error {
	return &ErrorCode{
		status: status,
		code:   code,
		msg:    fmt.Sprintf(msg, v...),
	}
}
