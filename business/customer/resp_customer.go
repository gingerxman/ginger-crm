package customer

import "github.com/gingerxman/ginger-crm/business/account"

type RCustomer struct {
	User *account.RUser `json:"user"`
}

type RLintCustomer struct {
	User *account.RUser `json:"user"`
	ConsumeCount int `json:"consume_count"`
	ConsumeMoney int `json:"consume_money"`
	LatestConsumeTime string `json:"latest_consume_time"`
}
