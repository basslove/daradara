package errors

import (
	"github.com/basslove/daradara/internal/api/pkg/customerror"
	stdErrors "github.com/cockroachdb/errors"
	"golang.org/x/xerrors"
	"net/http"
)

func ConvertError(err error) *Error {
	if err == nil {
		return nil
	}

	var appErr *Error
	if stdErrors.As(err, &appErr) {
		return appErr
	} else if customerror.IsServiceError(err) {
		return ConvertFromServiceError(customerror.ToServiceError(err))
	} else {
		return ErrInternalServerError.Wrap(err).(*Error)
	}
}

func ConvertFromServiceError(err *customerror.ServiceError) *Error {
	var appErr *Error

	switch err.Code {
	case customerror.WahahaError:
		appErr = ErrWahaha
	case customerror.GahahahaError:
		appErr = ErrGahaha
	case customerror.OhohoError:
		appErr = ErrOhoho
	default:
		appErr = NewError(http.StatusBadRequest, int(err.Code), err.Comment)
	}

	return appErr.Wrap(err).(*Error)
}

type Error struct {
	Code        int    `json:"code"`
	Status      int    `json:"status"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Detail      string `json:"detail"`
	TraceString string `json:"trace"`

	base  error
	err   error
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

func (e *Error) Is(errObj error) bool {
	var err *Error
	if stdErrors.As(errObj, &err) {
		return err.BaseError() == e
	}
	return false
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return e.Body
}

func (e *Error) BaseError() error {
	if e == nil {
		return nil
	}
	if e.base != nil {
		return e.base
	}
	return e
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}

func (e *Error) Wrap(errObj error) error {
	if e == nil {
		return nil
	}
	err := *e
	err.base = e
	err.err = errObj
	err.frame = xerrors.Caller(1)
	return &err
}

func (e *Error) Trace() error {
	if e == nil {
		return nil
	}

	err := *e
	err.base = e
	err.frame = xerrors.Caller(1)
	return &err
}
