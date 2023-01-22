package handler

import (
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/errors"
	"github.com/basslove/daradara/internal/api/infrastructure/api/middleware"
	"github.com/labstack/echo/v4"
	"time"
)

type CustomerAuth struct{}

func (cAuth *CustomerAuth) fetchFromContext(ctx echo.Context) (*model.Customer, error) {
	token := middleware.GetToken(ctx)
	if token == "" {
		return nil, errors.ErrNotFoundCustomerAccessToken
	}

	customer := middleware.GetCustomer(ctx)
	if customer == nil {
		return nil, errors.ErrIncorrectSession
	}

	//req := ctx.Request()
	//strCustomerID := strconv.Itoa(int(customer.ID))
	//sl := logger.With(req.Context(), zap.String("customer_id", strCustomerID))
	//
	//ctx.SetRequest(req.WithContext(logger.NewContext(req.Context(), sl)))

	now := time.Now()
	// 最終ログイン日から30日経過
	if customer.LastLoggedInAt.AddDate(0, 0, 30).Before(now) {
		return nil, errors.ErrRequiredReLogin
	}

	return customer, nil
}
