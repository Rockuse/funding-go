package campaign

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(campaign Campaign) (Campaign, error)
	FindAll() ([]Campaign, error)
	FindByUserId(userID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
func (r *repository) FindByUserId(userID int) ([]Campaign, error) {
	var campaignData []Campaign
	err := r.db.Find(&campaignData, "user_id=?", userID).Error
	if err != nil {
		return campaignData, err
	}
	return campaignData, nil
}
