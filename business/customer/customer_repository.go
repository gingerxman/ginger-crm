package customer

import (
	"context"
	"encoding/json"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business/account"
)

type CustomerRepository struct {
	eel.ServiceBase
}

func NewCustomerRepository(ctx context.Context) *CustomerRepository {
	service := new(CustomerRepository)
	service.Ctx = ctx
	return service
}

func (this *CustomerRepository) makeUser(userData interface{}) *account.User {
	userJson := userData.(map[string]interface{})
	id, _ := userJson["id"].(json.Number).Int64()
	user := account.NewUserFromOnlyId(this.Ctx, int(id))
	user.Name = userJson["name"].(string)
	user.Avatar = userJson["avatar"].(string)
	user.Sex = userJson["sex"].(string)
	user.Code = userJson["code"].(string)
	
	return user
}

func (this *CustomerRepository) makeCustomers(consumptionRecordDatas []interface{}) []*Customer {
	customers := make([]*Customer, 0)
	for _, recordData := range consumptionRecordDatas {
		recordData := recordData.(map[string]interface{})
		
		user := this.makeUser(recordData["user"])
		
		record := &consumptionRecord{}
		count, _ := recordData["consume_count"].(json.Number).Int64()
		record.ConsumeCount = int(count)
		money, _ := recordData["consume_money"].(json.Number).Int64()
		record.ConsumeMoney = int(money)
		record.LatestConsumeTime = recordData["latest_consume_time"].(string)
		
		customers = append(customers, &Customer{
			User: user,
			ConsumptionRecord: record,
		})
	}
	
	return customers
}

func (this *CustomerRepository) GetByConsumptionRecord() []*Customer {
	resp, err := eel.NewResource(this.Ctx).Get("ginger-mall", "consumption.user_consumption_records", eel.Map{
	})

	if err != nil {
		eel.Logger.Error(err)
		return nil
	}

	respData := resp.Data()
	recordDatas := respData.Get("records")
	return this.makeCustomers(recordDatas.MustArray())
}

func init() {
}
