package customer

import (
	"context"
	"encoding/json"
	"github.com/gingerxman/eel"
)

type FillCustomerService struct {
	eel.ServiceBase
}

func NewFillCustomerService(ctx context.Context) *FillCustomerService {
	service := new(FillCustomerService)
	service.Ctx = ctx
	return service
}

func (this *FillCustomerService) FillOne(customer *Customer, option eel.FillOption) {
	this.Fill([]*Customer{customer}, option)
}

func (this *FillCustomerService) Fill(customers []*Customer, option eel.FillOption)  {
	if len(customers) == 0 {
		return
	}
	
	ids := make([]int, 0)
	for _, customer := range customers {
		ids = append(ids, customer.Id)
	}
	if enableOption, ok := option["with_consumption_record"]; ok && enableOption {
		this.fillConsumptionRecord(customers, ids)
	}
}

func (this *FillCustomerService) fillConsumptionRecord(customers []*Customer, ids []int)  {
	userIds := make([]int, 0)
	user2customer := make(map[int]*Customer)
	for _, customer := range customers {
		userIds = append(userIds, customer.UserId)
		user2customer[customer.UserId] = customer
	}
	
	resp, err := eel.NewResource(this.Ctx).Get("ginger-order", "consumption.user_consumption_records", eel.Map{
		"user_ids": eel.ToJsonString(userIds),
	})
	
	if err != nil {
		eel.Logger.Error(err)
		return
	}
	
	respData := resp.Data()
	recordDatas := respData.Get("records").MustArray()
	for _, recordData := range recordDatas {
		recordData := recordData.(map[string]interface{})
		userId64, _ := recordData["consume_count"].(json.Number).Int64()
		userId := int(userId64)
		if customer, ok := user2customer[userId]; ok {
			record := &consumptionRecord{}
			count, _ := recordData["consume_count"].(json.Number).Int64()
			record.ConsumeCount = int(count)
			money, _ := recordData["consume_money"].(json.Number).Int64()
			record.ConsumeMoney = int(money)
			record.LatestConsumeTime = recordData["latest_consume_time"].(string)
			customer.ConsumptionRecord = record
		}
	}
}