package campaign

import (
	"errors"
	"fmt"
	"funding/src/app/module/user"
	"strconv"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	SaveCampaign(input CampaignInput) (Campaign, error)
	FindAll() ([]Campaign, error)
	FindByUserId(userId int) ([]Campaign, error)
	GetCampaignById(id CampaignUri) (Campaign, error)
	UpdateCampaign(input CampaignInput) (Campaign, error)
	UploadCampaignImage(images ImageInput, fileDir string) (CampaignImage, error)
}

type service struct {
	repository  Repository
	userService user.Service
}

func NewService(repository Repository, userService user.Service) *service {
	return &service{repository, userService}
}

func (s *service) SaveCampaign(input CampaignInput) (Campaign, error) {
	var data Campaign
	data.UserId = input.UserId
	data.Name = input.Name
	data.ShortDesc = input.ShortDesc
	data.Description = input.Description
	data.GoalAmmount = input.GoalAmmount
	data.CurrentAmmount = 0
	data.Perks = input.Perks
	data.CreatedDate = time.Now()
	data.CreatedBy = input.CreatedBy

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.Id)
	data.Slug = slug.Make(slugCandidate)
	inputUser, _ := s.userService.GetUserById(input.UserId)
	data.User = inputUser

	saved, err := s.repository.Save(data)
	if err != nil {
		return saved, err
	}
	return saved, nil
}

func (s *service) UpdateCampaign(input CampaignInput) (Campaign, error) {
	var data Campaign
	data.Id = input.Id
	data.UserId = input.UserId
	data.Name = input.Name
	data.ShortDesc = input.ShortDesc
	data.Description = input.Description
	data.GoalAmmount = input.GoalAmmount
	data.Perks = input.Perks
	data.ModifiedDate = time.Now()
	data.ModifiedBy = input.CreatedBy

	inputUser, _ := s.userService.GetUserById(input.UserId)
	data.User = inputUser
	saved, err := s.repository.Update(data)
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

func (s *service) UploadCampaignImage(input ImageInput, filePath string) (CampaignImage, error) {
	if input.IsPrimary {
		_, err := s.repository.MarkImagesNonPrimary(input.CampaignId)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	newImage := CampaignImage{}
	newImage.CampaignId = input.CampaignId
	newImage.FileName = filePath
	newImage.IsPrimary = input.IsPrimary
	newImage.CreatedDate = time.Now()
	newImage.CreatedBy = strconv.Itoa(input.UserId)

	result, err := s.repository.CreateImage(newImage)
	if err != nil {
		return result, err
	}
	return result, nil
}
