package customer

import (
	"context"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business/account"
)

type EncodeCustomerService struct {
	eel.ServiceBase
}

func NewEncodeCustomerService(ctx context.Context) *EncodeCustomerService {
	service := new(EncodeCustomerService)
	service.Ctx = ctx
	return service
}

func (this *EncodeCustomerService) Encode(customer *Customer) *RCustomer {
	return &RCustomer{
	}
}

func (this *EncodeCustomerService) EncodeMany(customers []*Customer) []*RCustomer {
	rows := make([]*RCustomer, len(customers))
	
	for i, customer := range customers {
		rows[i] = this.Encode(customer)
	}
	
	return rows
}

func (this *EncodeCustomerService) EncodeLintCustomer(customer *Customer) *RLintCustomer {
	var rUser *account.RUser
	if customer.User != nil {
		user := customer.User
		rUser = &account.RUser{
			Id: user.Id,
			Name: user.Name,
			Avatar: user.Avatar,
			Sex: user.Sex,
			Code: user.Code,
		}
	}
	
	consumptionRecord := customer.ConsumptionRecord
	return &RLintCustomer{
		User: rUser,
		ConsumeCount: consumptionRecord.ConsumeCount,
		ConsumeMoney: consumptionRecord.ConsumeMoney,
		LatestConsumeTime: consumptionRecord.LatestConsumeTime,
	}
}

func (this *EncodeCustomerService) EncodeManyLintCustomers(customers []*Customer) []*RLintCustomer {
	rows := make([]*RLintCustomer, len(customers))
	
	for i, customer := range customers {
		rows[i] = this.EncodeLintCustomer(customer)
	}
	
	return rows
}
