package transaction

import (
	util "funding/src/app/common/utilities"
	"time"
)

type Service interface {
	Add(transaction InputTransaction) (Transaction, error)
	GetById(data InputGetTransaction) (Transaction, []Transaction, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Add(input InputTransaction) (Transaction, error) {
	var data Transaction
	data.Id = util.Uuid()
	data.Code = input.Code
	data.CampaignId = input.CampaignId
	data.UserId = input.UserId
	data.Amount = input.Amount
	data.Status = input.Status
	data.ModifiedDate = time.Now()
	data.CreatedDate = time.Now()
	data.CreatedBy = input.UserId
	data.ModifiedBy = input.UserId
	err := s.repository.Add(data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *service) GetById(input InputGetTransaction) (Transaction, []Transaction, error) {
	if input.Type == "TRANSACTION" {
		var transaction Transaction
		transaction.Id = input.CampaignId
		transaction.UserId = input.UserId
		err := s.repository.GetOne(&transaction)
		if err != nil {
			return transaction, nil, err
		}
	} else if input.Type == "USER" {
		transaction, err := s.repository.GetByUser(input.UserId)
		if err != nil {
			return transaction, transaction, err
		}

	}

}
