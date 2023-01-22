package service

import (
	"context"
	"fmt"
	"github.com/basslove/daradara/internal/api/config"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/domain/repository"
	"github.com/basslove/daradara/internal/api/errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/xerrors"
)

type CustomerAuthenticator interface {
	FindOneByToken(context.Context, string) (*model.Customer, error)
}

type customerAuthenticatorImpl struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerAuthenticator(c repository.CustomerRepository) CustomerAuthenticator {
	return &customerAuthenticatorImpl{customerRepository: c}
}

func (s *customerAuthenticatorImpl) FindOneByToken(ctx context.Context, jwtToken string) (*model.Customer, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Get().Server.JwtSecret), nil
	})
	if err != nil {
		return nil, errors.ErrJwtExpired
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, xerrors.New("token error")
	}
	if _, ok = claims["sub"]; !ok {
		return nil, xerrors.New("sub does not exist")
	}

	email := fmt.Sprintf("%v", claims["sub"])
	customer, err := s.customerRepository.FindOneByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
