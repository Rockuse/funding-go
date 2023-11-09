package transaction

import "gorm.io/gorm"

type Repository interface {
	Add(transaction *Transaction) error
	GetOne(data *Transaction) error
	GetByUser(userId int) ([]Transaction, error)
	GetByCampaign(campaignId int) ([]Transaction, error)
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
	err := r.db.Joins("campaigns").Find(&transaction, "id=?", transaction.Id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetByUser(userId int) ([]Transaction, error) {
	var transaction []Transaction
	// err := r.db.Joins("join campaigns on campaigns.id = transactions.campaign_id").
	// 	Joins("left join campaign_images on campaign_images.campaign_id = campaigns.id").
	// 	Find(&transaction, "transactions.user_id=?", userId).Error
	err := r.db.Preload("Campaign.User").Find(&transaction, "transactions.user_id=?", userId).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) GetByCampaign(campaignId int) ([]Transaction, error) {
	var transaction []Transaction
	err := r.db.Preload("users").Where("campaign_id=?", campaignId).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil

}
