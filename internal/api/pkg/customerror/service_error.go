package customerror

import (
	"fmt"
	stdErrors "github.com/cockroachdb/errors"
)

type ServiceErrorCode int

const (
	NoErr         ServiceErrorCode = 0
	WahahaError   ServiceErrorCode = 215
	GahahahaError ServiceErrorCode = 216
	OhohoError    ServiceErrorCode = 217
	SystemError   ServiceErrorCode = 500
	Debug         ServiceErrorCode = 1000
	Unknown       ServiceErrorCode = 9999
)

type ServiceError struct {
	Code    ServiceErrorCode
	Comment string
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("Service error: %d - %s", e.Code, e.Comment)
}

func NewServiceError(code ServiceErrorCode, comment string) *ServiceError {
	return &ServiceError{
		Code:    code,
		Comment: comment,
	}
}

func IsServiceError(err error) bool {
	return ServiceErrCode(err) != Unknown
}

func ServiceErrCode(err error) ServiceErrorCode {
	if err != nil {
		return ToServiceError(err).Code
	}

	return Unknown
}

func ToServiceError(err error) *ServiceError {
	if err == nil {
		return nil
	}
	if base := new(ServiceError); stdErrors.As(err, &base) {
		return base
	}

	return &ServiceError{Code: Unknown}
}

func ServiceErrComment(err error) string {
	if err != nil {
		return ToServiceError(err).Comment
	}

	return ""
}
