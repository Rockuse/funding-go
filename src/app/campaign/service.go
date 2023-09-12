package campaign

import (
	"errors"
	"fmt"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	SaveCampaign(input CampaignInput) (Campaign, error)
	FindAll() ([]Campaign, error)
	FindByUserId(userId int) ([]Campaign, error)
	GetCampaignById(id CampaignUri) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SaveCampaign(input CampaignInput) (Campaign, error) {
	var data Campaign
	data.UserId = input.UserId
	data.Name = input.Name
	data.ShortDesc = input.ShortDesc
	data.Description = input.Description
	data.GoalAmmount = 0
	data.CurrentAmmount = 0
	data.Perks = "none"
	data.CreatedDate = time.Now()
	data.CreatedBy = input.CreatedBy
	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.Id)
	data.Slug = slug.Make(slugCandidate)
	inputUser, _ := 
	data.User = inputUser

	saved, err := s.repository.Save(data)
	if err != nil {
		return saved, err
	}
	return saved, nil
}

func (s *service) FindAll() ([]Campaign, error) {
	campaignList, err := s.repository.FindAll()
	if err != nil {
		return campaignList, err
	}
	if len(campaignList) == 0 {
		return campaignList, errors.New("not found")
	}
	return campaignList, nil
}

func (s *service) FindByUserId(userId int) ([]Campaign, error) {
	campaignData, err := s.repository.FindByUserId(userId)
	if err != nil {
		return campaignData, err
	}
	if len(campaignData) == 0 {
		return campaignData, errors.New("data not found")

	}
	return campaignData, nil
}

func (s *service) GetCampaignById(input CampaignUri) (Campaign, error) {
	campaignData, err := s.repository.FindById(input.ID)
	if err != nil {
		return campaignData, err
	}
	if campaignData.Id == 0 {
		return campaignData, errors.New("campaign not found")
	}
	return campaignData, nil
}
