package transaction

import (
	"funding/src/app/module/campaign"
	"time"
)

type Transaction struct {
	Id           int
	Code         string
	CampaignId   int
	UserId       int
	Amount       int
	Status       string
	ModifiedDate time.Time
	CreatedDate  time.Time
	CreatedBy    int
	ModifiedBy   int
	Campaign     campaign.Campaign
	Type         string
}
