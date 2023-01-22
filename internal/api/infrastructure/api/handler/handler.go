package handler

import (
	"github.com/basslove/daradara/internal/api/registry"
	usecase_customer "github.com/basslove/daradara/internal/api/usecase/customer"
	usecase_operator "github.com/basslove/daradara/internal/api/usecase/operator"
)

type Handler struct {
	CustomerPostCustomersHandler
	CustomerPostCustomersSignInHandler
	CustomerGetSightCategoriesHandler
	CustomerGetSightGenresHandler
	OperatorPostOperatorsHandler
	OperatorPostOperatorsSignInHandler
	OperatorGetSightCategoriesHandler
	OperatorGetSightGenresHandler
}

func NewHandler(repository *registry.Repository, service *registry.Service) (*Handler, error) {
	customerAuth := CustomerAuth{}
	operatorAuth := OperatorAuth{}

	return &Handler{
		CustomerPostCustomersHandler: CustomerPostCustomersHandler{
			InputPort: usecase_customer.NewPostCustomersUsecase(repository.CustomerRepository, repository.OperatorRepository),
		},
		CustomerPostCustomersSignInHandler: CustomerPostCustomersSignInHandler{
			InputPort: usecase_customer.NewPostCustomersSignInUsecase(repository.CustomerRepository, service.Throttler),
		},
		CustomerGetSightCategoriesHandler: CustomerGetSightCategoriesHandler{
			Auth:      customerAuth,
			InputPort: usecase_customer.NewGetSightCategoriesUsecase(repository.SightCategoryRepository),
		},
		CustomerGetSightGenresHandler: CustomerGetSightGenresHandler{
			Auth:      customerAuth,
			InputPort: usecase_customer.NewGetSightGenresUsecase(repository.SightGenreRelationQuery),
		},
		OperatorPostOperatorsHandler: OperatorPostOperatorsHandler{
			InputPort: usecase_operator.NewPostOperatorsUsecase(repository.CustomerRepository, repository.OperatorRepository),
		},
		OperatorPostOperatorsSignInHandler: OperatorPostOperatorsSignInHandler{
			InputPort: usecase_operator.NewPostOperatorsSignInUsecase(repository.OperatorRepository, service.Throttler),
		},
		OperatorGetSightCategoriesHandler: OperatorGetSightCategoriesHandler{
			Auth:      operatorAuth,
			InputPort: usecase_operator.NewGetSightCategoriesUsecase(repository.SightCategoryRepository),
		},
		OperatorGetSightGenresHandler: OperatorGetSightGenresHandler{
			Auth:      operatorAuth,
			InputPort: usecase_operator.NewGetSightGenresUsecase(repository.SightGenreRelationQuery),
		},
	}, nil
}
