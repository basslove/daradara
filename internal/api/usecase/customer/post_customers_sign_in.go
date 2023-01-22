package customer

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

type PostCustomersSignInInputPort interface {
	Execute(ctx context.Context, email, pass, IPAddress string) (*PostCustomersSignInOutput, error)
}

type PostCustomersSignInOutput struct {
	Customer *model.Customer
	Token    string
}

type PostCustomersSignInInteractor struct {
	customerRepository repository.CustomerRepository
	throttler          service.Throttler
}

func NewPostCustomersSignInUsecase(c repository.CustomerRepository, t service.Throttler) PostCustomersSignInInputPort {
	return &PostCustomersSignInInteractor{customerRepository: c, throttler: t}
}

func (i *PostCustomersSignInInteractor) Execute(ctx context.Context, email, pass, IPAddress string) (*PostCustomersSignInOutput, error) {
	isBlocked, err := i.throttler.IsBlocked(ctx, IPAddress, model.ThrottleKeyTypeIP)
	if err != nil {
		return nil, err
	}
	if isBlocked {
		return nil, errors.ErrUnauthorized
	}

	customer, err := i.customerRepository.FindOneByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		if err = i.throttler.Increase(ctx, IPAddress, model.ThrottleKeyTypeIP); err != nil {
			return nil, err
		}
		return nil, errors.ErrNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(customer.CryptedPassword), []byte(pass))
	if err != nil {
		if err = i.throttler.Increase(ctx, IPAddress, model.ThrottleKeyTypeIP); err != nil {
			return nil, err
		}
		return nil, errors.ErrBadRequest
	}

	customer.LastLoggedInAt = time.Now()
	customer.LastAccessedAt = time.Now()
	if err = i.customerRepository.Update(ctx, customer); err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 有効期限
	claims["iat"] = time.Now().Unix()                     // 発行時間
	claims["sub"] = customer.Email                        // email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtString, err := token.SignedString([]byte(config.Get().Server.JwtSecret))
	if err != nil {
		return nil, err
	}

	return &PostCustomersSignInOutput{Customer: customer, Token: jwtString}, nil
}
