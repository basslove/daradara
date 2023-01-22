package errors

import "net/http"

var (
	ErrBadRequest          = NewError(http.StatusBadRequest, 400, "bad request")
	ErrUnauthorized        = NewError(http.StatusUnauthorized, 401, "not login")
	ErrForbidden           = NewError(http.StatusForbidden, 403, "forbidden")
	ErrNotFound            = NewError(http.StatusNotFound, 404, "not found")
	ErrInternalServerError = NewError(http.StatusInternalServerError, 500, "internal server error")
)
