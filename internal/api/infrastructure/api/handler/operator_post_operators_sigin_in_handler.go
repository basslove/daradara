package handler

import (
	"github.com/basslove/daradara/internal/api/errors"
	"github.com/basslove/daradara/internal/api/infrastructure/api/openapi_service"
	usecase "github.com/basslove/daradara/internal/api/usecase/operator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type OperatorPostOperatorsSignInHandler struct {
	InputPort usecase.PostOperatorsSignInInputPort
}

func (h *OperatorPostOperatorsSignInHandler) OperatorPostOperatorsSignIn(ctx echo.Context) error {
	var request openapi_service.OperatorPostOperatorsSignInRequestBody
	if err := ctx.Bind(&request); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	output, err := h.InputPort.Execute(ctx.Request().Context(), request.Email, request.Password, ctx.RealIP())
	if err != nil {
		return err
	}

	var operator *openapi_service.Operator
	if output.Operator != nil {
		operator = openapi_service.BuildOperatorPostOperatorsSignInResponse(output.Operator)
	}

	return ctx.JSON(http.StatusCreated, openapi_service.OperatorPostOperatorsSignInResponse{Operator: operator, Token: &output.Token})
}
