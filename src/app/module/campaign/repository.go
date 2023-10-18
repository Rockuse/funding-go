package campaign

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	FindAll() ([]Campaign, error)
	FindByUserId(userID int) ([]Campaign, error)
	FindById(ID int) (Campaign, error)
	CreateImage(campaignImage CampaignImage) (CampaignImage, error)
	MarkImagesNonPrimary(campaignId int) (bool, error)
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

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	data := r.db.Model(&campaign).Where("user_id=? and id=?", campaign.UserId, campaign.Id).Updates(campaign) //Save(&campaign).Where("user_id=?", campaign.UserId).Error
	if data.Error != nil {
		return campaign, data.Error
	}
	if data.RowsAffected == 0 {
		return campaign, errors.New("no update")
	}
	return campaign, nil
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary =true").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
func (r *repository) FindByUserId(userID int) ([]Campaign, error) {
	var campaignData []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary =true").Find(&campaignData, "user_id=?", userID).Error
	if err != nil {
		return campaignData, err
	}
	return campaignData, nil
}

func (r *repository) FindById(ID int) (Campaign, error) {
	var campaignData Campaign
	err := r.db.Preload("User").Preload("CampaignImages").Find(&campaignData, "id=?", ID).Error
	if err != nil {
		return campaignData, err
	}
	return campaignData, nil
}

func (r *repository) CreateImage(campaignImage CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

func (r *repository) UpdatePrimary(campaignId int) (bool, error) {
	data := r.db.Model(&CampaignImage{}).Update("is_primary", false).Where("campaign_id", campaignId)

	if data.Error != nil {
		return false, data.Error
	} else if data.RowsAffected == 0 {
		return false, errors.New("data not found")
	}
	return true, nil
}

func (r *repository) MarkImagesNonPrimary(id int) (bool, error) {
	data := r.db.Model(&CampaignImage{}).Where("campaign_id", id).Update("is_primary", false)
	if data.Error != nil {
		return false, data.Error
	}
	return true, nil
}
