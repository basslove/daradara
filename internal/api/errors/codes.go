package errors

import (
	"github.com/basslove/daradara/internal/api/pkg/customerror"
	"net/http"
)

var (
	// common
	ErrBadRequest          = NewError(http.StatusBadRequest, 400, "bad request")
	ErrUnauthorized        = NewError(http.StatusUnauthorized, 401, "not login")
	ErrForbidden           = NewError(http.StatusForbidden, 403, "forbidden")
	ErrNotFound            = NewError(http.StatusNotFound, 404, "not found")
	ErrAlreadyExists       = NewError(http.StatusConflict, 409, "already exists")
	ErrInternalServerError = NewError(http.StatusInternalServerError, 500, "internal server error")

	// NOT FOUND(404XX)
	ErrNotFoundCustomerAccessToken = NewError(http.StatusNotFound, 40401, "customer access token not found")

	// Unauthorized(401XX)
	ErrIncorrectSession      = NewError(http.StatusUnauthorized, 40101, "customer session is incorrect")
	ErrRequiredReLogin       = NewError(http.StatusUnauthorized, 40102, "customer is re login")
	ErrCustomerNilNotAllowed = NewError(http.StatusUnauthorized, 40103, "customer is nil")
	ErrOperatorNilNotAllowed = NewError(http.StatusUnauthorized, 40104, "operator is nil")
	ErrJwtExpired            = NewError(http.StatusUnauthorized, 40105, "jwt token is expired")

	// custom
	ErrWahaha = NewError(http.StatusBadRequest, int(customerror.WahahaError), "this is wahaha error")
	ErrGahaha = NewError(http.StatusBadRequest, int(customerror.GahahahaError), "this is gahaha error")
	ErrOhoho  = NewError(http.StatusBadRequest, int(customerror.OhohoError), "this is ohoho error")
)
