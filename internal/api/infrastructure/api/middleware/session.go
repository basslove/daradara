package middleware

import (
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
	"strings"
)

const (
	authCustomerKey = "daradara:customerkey"
	authOperatorKey = "daradara:operatorkey"
	authScheme      = "Bearer"
)

func GetToken(ctx echo.Context) string {
	token, err := getTokenFromHeader(ctx)
	if err != nil {
		return ""
	}
	return token
}

func getTokenFromHeader(c echo.Context) (string, error) {
	authorization := c.Request().Header.Get(echo.HeaderAuthorization)

	l := len(authScheme)
	if len(authorization) > l+1 && authorization[:l] == authScheme {
		return strings.TrimSpace(authorization[l+1:]), nil
	}
	return "", xerrors.New("missing token")
}

func GetCustomer(ctx echo.Context) *model.Customer {
	customer := ctx.Get(authCustomerKey)

	if v, ok := customer.(*model.Customer); ok {
		return v
	}
	return nil
}
func SetCustomer(ctx echo.Context, customer *model.Customer) {
	ctx.Set(authCustomerKey, customer)
}

func GetOperator(ctx echo.Context) *model.Operator {
	operator := ctx.Get(authOperatorKey)

	if v, ok := operator.(*model.Operator); ok {
		return v
	}
	return nil
}
func SetOperator(ctx echo.Context, operator *model.Operator) {
	ctx.Set(authOperatorKey, operator)
}
