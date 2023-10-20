package transaction

import "gorm.io/gorm"

type Repository interface {
	Add(transaction Transaction) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Add(transaction *Transaction) error {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}
