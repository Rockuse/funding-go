package transaction

type InputTransaction struct {
	CampaignId int `json:"campaign_id" binding:"required"`
	UserId     int `json:"user_id"`
	Amount     int `json:"amount" binding:"required"`
	Status     int `json:"status"`
}
