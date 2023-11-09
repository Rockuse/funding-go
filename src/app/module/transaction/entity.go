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
	Status       int
	Campaign     campaign.Campaign
	ModifiedDate time.Time
	CreatedDate  time.Time
	CreatedBy    int
	ModifiedBy   int
}
