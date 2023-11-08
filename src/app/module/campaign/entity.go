package campaign

import (
	"funding/src/app/module/user"
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
	CreatedBy      int
	ModifiedBy     int
	CampaignImages []CampaignImage
	User           user.User
}

type CampaignImage struct {
	Id           int
	CampaignId   int    `json:"campaign_id" binding:"required"`
	FileName     string `json:"file_name" binding:"required"`
	IsPrimary    bool   `json:"is_primary" binding:"required"`
	ModifiedDate time.Time
	CreatedDate  time.Time
	CreatedBy    int
	ModifiedBy   int
}

type CampaignUri struct {
	ID int `uri:"id" binding:"required"`
}
