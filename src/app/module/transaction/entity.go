package transaction

import "time"

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
}
