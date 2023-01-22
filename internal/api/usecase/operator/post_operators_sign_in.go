package operator

import (
	"context"
	"github.com/basslove/daradara/internal/api/config"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/domain/repository"
	"github.com/basslove/daradara/internal/api/domain/service"
	"github.com/basslove/daradara/internal/api/errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type PostOperatorsSignInInputPort interface {
	Execute(ctx context.Context, email, pass, IPAddress string) (*PostOperatorsSignInOutput, error)
}

type PostOperatorsSignInOutput struct {
	Operator *model.Operator
	Token    string
}

type PostOperatorsSignInInteractor struct {
	operatorRepository repository.OperatorRepository
	throttler          service.Throttler
}

func NewPostOperatorsSignInUsecase(c repository.OperatorRepository, t service.Throttler) PostOperatorsSignInInputPort {
	return &PostOperatorsSignInInteractor{operatorRepository: c, throttler: t}
}

func (i *PostOperatorsSignInInteractor) Execute(ctx context.Context, email, pass, IPAddress string) (*PostOperatorsSignInOutput, error) {
	isBlocked, err := i.throttler.IsBlocked(ctx, IPAddress, model.ThrottleKeyTypeIP)
	if err != nil {
		return nil, err
	}
	if isBlocked {
		return nil, errors.ErrUnauthorized
	}

	operator, err := i.operatorRepository.FindOneByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if operator == nil {
		if err = i.throttler.Increase(ctx, IPAddress, model.ThrottleKeyTypeIP); err != nil {
			return nil, err
		}
		return nil, errors.ErrNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(operator.CryptedPassword), []byte(pass))
	if err != nil {
		if err = i.throttler.Increase(ctx, IPAddress, model.ThrottleKeyTypeIP); err != nil {
			return nil, err
		}
		return nil, errors.ErrBadRequest
	}

	operator.LastLoggedInAt = time.Now()
	operator.LastAccessedAt = time.Now()
	if err = i.operatorRepository.Update(ctx, operator); err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 有効期限
	claims["iat"] = time.Now().Unix()                     // 発行時間
	claims["sub"] = operator.Email                        // email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtString, err := token.SignedString([]byte(config.Get().Server.JwtSecret))
	if err != nil {
		return nil, err
	}

	return &PostOperatorsSignInOutput{Operator: operator, Token: jwtString}, nil
}
