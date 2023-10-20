package transaction

type Service interface {
	Add(transaction Transaction) (Transaction, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Add(transaction Transaction) (Transaction, error) {
	var data Transaction
	
}
