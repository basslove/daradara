package operator

import (
	"context"
	"github.com/basslove/daradara/internal/api/domain/model"
	"github.com/basslove/daradara/internal/api/domain/repository"
	"github.com/basslove/daradara/internal/api/errors"
	"github.com/basslove/daradara/internal/api/interface/presenter/forms"
	"golang.org/x/crypto/bcrypt"
)

type PostOperatorsInputPort interface {
	Execute(ctx context.Context, fm *forms.OperatorPostOperatorsForm) (*PostOperatorsOutput, error)
}

type PostOperatorsOutput struct {
	Operator *model.Operator
}

type PostOperatorsInteractor struct {
	customerOperator   repository.CustomerRepository
	operatorRepository repository.OperatorRepository
}

func NewPostOperatorsUsecase(c repository.CustomerRepository, o repository.OperatorRepository) PostOperatorsInputPort {
	return &PostOperatorsInteractor{customerOperator: c, operatorRepository: o}
}

func (i *PostOperatorsInteractor) Execute(ctx context.Context, fm *forms.OperatorPostOperatorsForm) (*PostOperatorsOutput, error) {
	if fm == nil {
		return nil, errors.ErrInternalServerError
	}
	customer, err := i.customerOperator.FindOneByEmail(ctx, fm.Email)
	if err != nil {
		return nil, err
	}
	if customer != nil {
		return nil, errors.ErrAlreadyExists
	}

	operator, err := i.operatorRepository.FindOneByEmail(ctx, fm.Email)
	if err != nil {
		return nil, err
	}
	if operator != nil {
		return nil, errors.ErrAlreadyExists
	}

	if fm.Password != fm.PasswordConfirmation {
		return nil, errors.ErrBadRequest
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(fm.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.ErrInternalServerError.Wrap(err)
	}

	newOperator := &model.Operator{
		Email:           fm.Email,
		CryptedPassword: string(hashPass),
		DisplayName:     fm.DisplayName,
		Name:            fm.Name,
		Belong:          fm.Belong,
		Level:           fm.Level,
		IsGod:           fm.IsGod,
		IsValid:         true,
	}
	newID, err := i.operatorRepository.Create(ctx, newOperator)
	if err != nil {
		return nil, err
	}
	newOperator.ID = uint64(newID)

	return &PostOperatorsOutput{Operator: newOperator}, nil
}
