package campaign

import (
	"funding/src/app/user"
	"time"
)

type Campaign struct {
	Id             int
	UserId         int
	Name           string
	ShortDesc      string
	Description    string
	GoalAmmount    int
	CurrentAmmount int
	Perks          string
	BackerCount    int
	Slug           string
	ModifiedDate   time.Time
	CreatedDate    time.Time
	CreatedBy      string
	ModifiedBy     string
	CampaignImages []CampaignImage
	User           user.User
}

type CampaignImage struct {
	Id           int
	CampaignId   int
	FileName     string
	IsPrimary    bool
	ModifiedDate time.Time
	CreatedDate  time.Time
	CreatedBy    string
	ModifiedBy   string
}

type CampaignInput struct {
	UserId      int    `json:"userid"`
	Name        string `json:"name" binding:"required"`
	ShortDesc   string `json:"shortdesc" binding:"required"`
	Description string `json:"description"`
	CreatedBy   string
}
