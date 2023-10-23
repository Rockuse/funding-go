package transaction

import "gorm.io/gorm"

type Repository interface {
	Add(transaction Transaction) error
	GetOne(data *Transaction) error
	GetByUser(userId int) ([]Transaction, error)
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

func (r *repository) GetOne(transaction *Transaction) error {
	err := r.db.Find(&transaction, "id=?", transaction.Id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetByUser(userId int) ([]Transaction, error) {
	var transaction []Transaction
	err := r.db.InnerJoins("Campaign").Find(&transaction, "user_id=?", userId).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
