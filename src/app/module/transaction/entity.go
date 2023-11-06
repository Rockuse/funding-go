package transaction

import (
	"funding/src/app/module/campaign"
	"funding/src/app/module/user"
	"time"
)

type Transaction struct {
	Id           int
	Code         string
	CampaignId   int
	UserId       int
	Amount       int
	Status       string
	Campaign     campaign.Campaign
	User         user.User
	ModifiedDate time.Time
	CreatedDate  time.Time
	CreatedBy    int
	ModifiedBy   int
}
