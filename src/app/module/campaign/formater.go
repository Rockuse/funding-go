package campaign

import (
	"strings"
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
	CreatedBy      int       `json:"created_by"`
}

type CampaignDetailFormat struct {
	Detail CampaignFormat   `json:"detail"`
	User   CampaignUser     `json:"user"`
	Perk   []string         `json:"perk"`
	Images []CampaignImages `json:"images"`
}

type CampaignUser struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}
type CampaignImages struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
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
		formater.ImageURL = host + "/images" + data.CampaignImages[0].FileName
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

func FormatDetail(data Campaign, host string) CampaignDetailFormat {
	formater := CampaignDetailFormat{
		Detail: CampaignFormat{
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
		},
		// User:   CampaignUser{},
		// Perk:   []string,
		// Images: []CampaignImages{},

	}
	formater.Perk = strings.Split(data.Perks, ",")
	formater.User.ImageUrl = host + "/images" + data.User.Avatar_file_name
	formater.User.Name = data.User.Name

	var images []CampaignImages
	for _, image := range data.CampaignImages {
		images = append(images, FormatImage(image, host))
	}
	formater.Images = images
	return formater
}

func FormatImage(data CampaignImage, host string) CampaignImages {
	var campaignImage = CampaignImages{
		ImageUrl:  host + "/images" + data.FileName,
		IsPrimary: data.IsPrimary,
	}

	return campaignImage
}
func (p CampaignInput) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":         p.Name,
		"shortdesc":    p.ShortDesc,
		"description":  p.Description,
		"goal_ammount": p.GoalAmmount,
		"perks":        p.Perks,
	}
}
