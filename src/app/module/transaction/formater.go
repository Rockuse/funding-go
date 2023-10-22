package transaction

type TransactionFormat struct {
	Id           int
	Code         string
	CampaignId   int
	UserId       int
	Amount       int
	Status       string
}

func Format(data Transaction) TransactionFormat {
	return TransactionFormat{}
}
