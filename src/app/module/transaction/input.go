package transaction

type InputTransaction struct {
	Code       string
	CampaignId int
	UserId     int
	Amount     int
	Status     string
}

type InputGetTransaction struct {
	UserId     int
	CampaignId int
	Type       string
}
	