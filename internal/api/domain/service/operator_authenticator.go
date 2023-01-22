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

type OperatorAuthenticator interface {
	FindOneByToken(context.Context, string) (*model.Operator, error)
}

type operatorAuthenticatorImpl struct {
	operatorRepository repository.OperatorRepository
}

func NewOperatorAuthenticator(c repository.OperatorRepository) OperatorAuthenticator {
	return &operatorAuthenticatorImpl{operatorRepository: c}
}

func (s *operatorAuthenticatorImpl) FindOneByToken(ctx context.Context, jwtToken string) (*model.Operator, error) {
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
	operator, err := s.operatorRepository.FindOneByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return operator, nil
}
