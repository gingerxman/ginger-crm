package customer

import (
	"context"
	"github.com/gingerxman/eel"
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
	consumptionRecord := customer.ConsumptionRecord
	return &RCustomer{
		Id: customer.Id,
		UserId: customer.UserId,
		Unionid: customer.Unionid,
		Code: customer.Code,
		
		Sex: customer.Sex,
		Name: customer.Name,
		Avatar: customer.Avatar,
		Source: customer.Source,
		
		ConsumeCount: consumptionRecord.ConsumeCount,
		ConsumeMoney: consumptionRecord.ConsumeMoney,
		LatestConsumeTime: consumptionRecord.LatestConsumeTime,
	}
}

func (this *EncodeCustomerService) EncodeMany(customers []*Customer) []*RCustomer {
	rows := make([]*RCustomer, len(customers))
	
	for i, customer := range customers {
		rows[i] = this.Encode(customer)
	}
	
	return rows
}