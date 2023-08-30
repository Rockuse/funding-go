package campaign

type Service interface {
	SaveCampaign(input CampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SaveCampaign(input CampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDesc = input.ShortDesc
	campaign.Description = input.Description
	campaign.GoalAmmount = 0
	campaign.CurrentAmmount = 0
	campaign.Perks = "none"

	saved, err := s.repository.Save(campaign)
	if err != nil {
		return saved, err
	}
	return saved, nil
}
