package handler

import (
	"fmt"
	"github.com/basslove/daradara/internal/api/config"
	"github.com/basslove/daradara/internal/api/errors"
	"github.com/basslove/daradara/internal/api/logger"
	stdErrors "github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

func HandlerJSONError(err error, ctx echo.Context) {
	if ctx.Response().Committed {
		return
	}

	var appErr *errors.Error
	logged := true

	if httpErr, ok := err.(*echo.HTTPError); ok {
		// std error(echo)
		appErr = convertError(httpErr.Code, err)
		logged = false
	} else {
		// custom error
		appErr = err.(*errors.Error)
	}

	if logged {
		if appErr.Status >= 500 {
			logger.Error(ctx.Request().Context(), appErr)
		} else {
			logger.Infow(ctx.Request().Context(), appErr.Error(), zap.String("stacktrace", fmt.Sprintf("%+v", appErr)))
		}
	}

	if err := responseJson(appErr, ctx); err != nil {
		logger.Error(ctx.Request().Context(), err)
	}
}

func responseJson(e *errors.Error, ctx echo.Context) error {
	if config.Get().Server.Debug {
		e.TraceString = fmt.Sprintf("%+v\n", e)
		e.Detail = e.Error()

		for err := stdErrors.Unwrap(e); err != nil; err = stdErrors.Unwrap(err) {
			e.Detail = fmt.Sprintf("%s > %s", e.Detail, err.Error())
		}
	}
	return ctx.JSON(e.Status, e)
}

func convertError(statusCode int, err error) *errors.Error {
	switch statusCode {
	case http.StatusBadRequest:
		return errors.ErrBadRequest.Wrap(err).(*errors.Error)
	case http.StatusUnauthorized:
		return errors.ErrUnauthorized.Wrap(err).(*errors.Error)
	case http.StatusForbidden:
		return errors.ErrForbidden.Wrap(err).(*errors.Error)
	case http.StatusNotFound:
		return errors.ErrNotFound.Wrap(err).(*errors.Error)
	case http.StatusConflict:
		return errors.ErrAlreadyExists.Wrap(err).(*errors.Error)
	default:
		return errors.ErrInternalServerError.Wrap(err).(*errors.Error)
	}
}
