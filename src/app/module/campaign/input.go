package campaign

import "funding/src/app/module/user"

type ImageInput struct {
	CampaignId int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary" binding:"required"`
	UserId     int
}

type CampaignInput struct {
	Id          int
	UserId      int    `json:"user_id"`
	Name        string `json:"name" binding:"required"`
	ShortDesc   string `json:"shortdesc" binding:"required"`
	Description string `json:"description"`
	GoalAmmount int    `json:"goal_ammount"`
	Perks       string `json:"perks"`
	CreatedBy   int
	User        user.User
}
