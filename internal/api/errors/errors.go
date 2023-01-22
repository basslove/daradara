package errors

import "golang.org/x/xerrors"

type Error struct {
	Code        int    `json:"code"`
	Status      int    `json:"status"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Detail      string `json:"detail"`
	TraceString string `json:"trace"`

	frame xerrors.Frame
}

func NewError(status, code int, message string) *Error {
	return &Error{
		Code:   code,
		Status: status,
		Body:   message,
		frame:  xerrors.Caller(1),
	}
}
