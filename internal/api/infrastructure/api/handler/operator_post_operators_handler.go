package handler

import (
	"github.com/basslove/daradara/internal/api/errors"
	"github.com/basslove/daradara/internal/api/infrastructure/api/openapi_service"
	"github.com/basslove/daradara/internal/api/interface/presenter/forms"
	usecase "github.com/basslove/daradara/internal/api/usecase/operator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type OperatorPostOperatorsHandler struct {
	InputPort usecase.PostOperatorsInputPort
}

func (h *OperatorPostOperatorsHandler) OperatorPostOperators(ctx echo.Context) error {
	var request openapi_service.OperatorPostOperatorsRequestBody
	if err := ctx.Bind(&request); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	requestForm := forms.NewOperatorPostOperatorsForm(request)
	if err := requestForm.Validate(); err != nil {
		return errors.ErrBadRequest.Wrap(err)
	}

	output, err := h.InputPort.Execute(ctx.Request().Context(), requestForm)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, openapi_service.OperatorPostOperatorsResponse{Id: output.Operator.ID})
}
