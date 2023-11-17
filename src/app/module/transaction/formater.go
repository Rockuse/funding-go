package transaction

import (
	"sort"
	"time"
)

type TransactionFormat struct {
	Id           int       `json:"id"`
	Code         string    `json:"code"`
	CampaignName string    `json:"campaign_name"`
	UserId       string    `json:"user_name"`
	Amount       int       `json:"amount"`
	CreatedDate  time.Time `json:"created_date"`
	Status       int       `json:"status"`
}

func FormatTransaction(data Transaction) TransactionFormat {
	var format TransactionFormat
	format.Id = data.Id
	format.Code = data.Code
	format.CampaignName = data.Campaign.Name
	format.UserId = data.Campaign.User.Name
	format.Amount = data.Amount
	format.CreatedDate = data.CreatedDate
	format.Status = data.Status
	return format
}

func FormatListTransaction(data []Transaction) []TransactionFormat {
	sort.Slice(data, func(i, j int) bool {
		return data[i].CreatedDate.After(data[j].CreatedDate)
	})
	if len(data) == 0 {
		return []TransactionFormat{}
	}
	var format []TransactionFormat
	for _, trans := range data {
		format = append(format, FormatTransaction(trans))
	}

	return format
}
