package registry

type Service struct{}

func NewService(repository *Repository) (*Service, error) {
	s := Service{}

	return &s, nil
}
