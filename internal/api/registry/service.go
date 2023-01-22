package registry

import "github.com/basslove/daradara/internal/api/domain/service"

type Service struct {
	Throttler             service.Throttler
	CustomerAuthenticator service.CustomerAuthenticator
	OperatorAuthenticator service.OperatorAuthenticator
}

func NewService(repository *Repository) (*Service, error) {
	throttler := service.NewThrottler(repository.ThrottleRepository)
	customerAuthenticator := service.NewCustomerAuthenticator(repository.CustomerRepository)
	operatorAuthenticator := service.NewOperatorAuthenticator(repository.OperatorRepository)

	s := Service{
		Throttler:             throttler,
		CustomerAuthenticator: customerAuthenticator,
		OperatorAuthenticator: operatorAuthenticator,
	}

	return &s, nil
}
