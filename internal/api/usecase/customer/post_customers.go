package customer

import (
	"context"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/domain/repository"
	"github.com/basslove/daradara/internal/api/errors"
	"github.com/basslove/daradara/internal/api/interface/presenter/forms"
	"golang.org/x/crypto/bcrypt"
)

type PostCustomersInputPort interface {
	Execute(ctx context.Context, fm *forms.CustomerPostCustomersForm) (*PostCustomersOutput, error)
}

type PostCustomersOutput struct {
	Customer *model.Customer
}

type PostCustomersInteractor struct {
	operatorRepository repository.OperatorRepository
	customerRepository repository.CustomerRepository
}

func NewPostCustomersUsecase(c repository.CustomerRepository, o repository.OperatorRepository) PostCustomersInputPort {
	return &PostCustomersInteractor{customerRepository: c, operatorRepository: o}
}

func (i *PostCustomersInteractor) Execute(ctx context.Context, fm *forms.CustomerPostCustomersForm) (*PostCustomersOutput, error) {
	if fm == nil {
		return nil, errors.ErrInternalServerError
	}
	operator, err := i.operatorRepository.FindOneByEmail(ctx, fm.Email)
	if err != nil {
		return nil, errors.ConvertError(err).Wrap(err)
	}
	if operator != nil {
		return nil, errors.ErrAlreadyExists
	}

	customer, err := i.customerRepository.FindOneByEmail(ctx, fm.Email)
	if err != nil {
		return nil, err
	}
	if customer != nil {
		return nil, errors.ErrAlreadyExists
	}

	if fm.Password != fm.PasswordConfirmation {
		return nil, errors.ErrBadRequest
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(fm.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.ErrInternalServerError.Wrap(err)
	}

	newCustomer := &model.Customer{
		Email:              fm.Email,
		CryptedPassword:    string(hashPass),
		Introduction:       fm.Introduction,
		PhoneNumber:        fm.PhoneNumber,
		DisplayName:        fm.DisplayName,
		Name:               fm.Name,
		Birthday:           fm.Birthday,
		Generation:         fm.Generation,
		Gender:             fm.Gender,
		AllowPlanDisplayed: fm.AllowPlansDisplayed,
		IsValid:            true,
	}
	newID, err := i.customerRepository.Create(ctx, newCustomer)
	if err != nil {
		return nil, err
	}
	newCustomer.ID = uint64(newID)

	return &PostCustomersOutput{Customer: newCustomer}, nil
}
