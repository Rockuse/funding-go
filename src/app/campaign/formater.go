package campaign

import (
	"strconv"
	"time"
)

type CampaignFormat struct {
	Id             int       `json:"id"`
	UserId         int       `json:"user_id"`
	Name           string    `json:"name"`
	ShortDesc      string    `json:"short_description"`
	ImageURL       string    `json:"image_url"`
	GoalAmmount    int       `json:"goal_ammount"`
	CurrentAmmount int       `json:"current_ammount"`
	Slug           string    `json:"slug"`
	CreatedDate    time.Time `json:"created_date"`
	CreatedBy      string    `json:"create_dby"`
}

func FormatCampaign(data Campaign, host string) CampaignFormat {
	formater := CampaignFormat{
		Id:             data.Id,
		UserId:         data.UserId,
		Name:           data.Name,
		ShortDesc:      data.ShortDesc,
		ImageURL:       "",
		GoalAmmount:    data.GoalAmmount,
		CurrentAmmount: data.CurrentAmmount,
		Slug:           data.Slug,
		CreatedDate:    data.CreatedDate,
		CreatedBy:      data.CreatedBy,
	}
	if len(data.CampaignImages) > 0 {
		formater.ImageURL = host + "/images/" + strconv.Itoa(data.CampaignImages[0].Id)
	}

	return formater
}

func FormatAllCampaigns(data []Campaign, host string) []CampaignFormat {
	if len(data) == 0 {
		return []CampaignFormat{}

	}
	formater := []CampaignFormat{}
	for _, campaign := range data {
		formater = append(formater, FormatCampaign(campaign, host))
	}
	return formater
}
