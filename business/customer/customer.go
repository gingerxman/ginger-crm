package customer

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business/account"
)

type consumptionRecord struct {
	ConsumeCount int
	ConsumeMoney int
	LatestConsumeTime string
}

type Customer struct {
	eel.EntityBase
	Id                int
	
	User *account.User
	ConsumptionRecord *consumptionRecord
}

func init() {
}
