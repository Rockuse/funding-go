package campaign

import "time"

type CampaignFormat struct {
	Id             int       `json:"id"`
	UserId         int       `json:"userid"`
	Name           string    `json:"name"`
	ShortDesc      string    `json:"shortdesc"`
	Description    string    `json:"description"`
	GoalAmmount    int       `json:"goalammount"`
	CurrentAmmount int       `json:"currentammount"`
	Perks          string    `json:"perks"`
	BackerCount    int       `json:"backercount"`
	Slug           string    `json:"slug"`
	CreatedDate    time.Time `json:"createddate"`
	CreatedBy      string    `json:"createdby"`
}

func FormatCampaign(data Campaign) CampaignFormat {
	formater := CampaignFormat{
		Id:             data.Id,
		UserId:         data.UserId,
		Name:           data.Name,
		ShortDesc:      data.ShortDesc,
		Description:    data.Description,
		GoalAmmount:    data.GoalAmmount,
		CurrentAmmount: data.CurrentAmmount,
		Perks:          data.Perks,
		BackerCount:    data.BackerCount,
		Slug:           data.Slug,
		CreatedDate:    data.CreatedDate,
		CreatedBy:      data.CreatedBy,
	}
	return formater
}
