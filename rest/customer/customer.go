package customer

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business/account"
	b_customer "github.com/gingerxman/ginger-crm/business/customer"
)

type Customer struct {
	eel.RestResource
}

func (this *Customer) Resource() string {
	return "customer.customer"
}

func (this *Customer) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"?id:int"},
		"PUT": []string{},
	}
}

func (this *Customer) Get(ctx *eel.Context) {
	req := ctx.Request
	bCtx := ctx.GetBusinessContext()
	
	corp := account.GetCorpFromContext(bCtx)
	var customer *b_customer.Customer
	id, _ := req.GetInt("id", 0)
	if id == 0 {
		user := account.GetUserFromContext(bCtx)
		customer = b_customer.NewCustomerRepository(bCtx).GetCustomerByUserIdInCorp(corp, user)
	} else {
		customer = b_customer.NewCustomerRepository(bCtx).GetCustomerByIdInCorp(corp, id)
	}
	
	if customer == nil {
		ctx.Response.Error("customer:invalid_customer", "")
		return
	}
	
	b_customer.NewFillCustomerService(bCtx).FillOne(customer, eel.FillOption{
		"with_consumption_record": true,
	})
	data := b_customer.NewEncodeCustomerService(bCtx).Encode(customer)
	
	ctx.Response.JSON(data)
}

func (this *Customer) Put(ctx *eel.Context) {
	bCtx := ctx.GetBusinessContext()
	
	user := account.GetUserFromContext(bCtx)
	if user.Id == 0 {
		ctx.Response.Error("customer:invalid_user_id", "user_id(0)")
		return
	}
	
	corp := account.GetCorpFromContext(bCtx)
	if corp.Id == 0 {
		ctx.Response.Error("customer:invalid_corp", "corp_id(0)")
		return
	}
	
	customer, err := b_customer.NewCustomerForCorp(bCtx, user, corp)
	if err != nil {
		ctx.Response.Error("customer:create_user_fail", err.Error())
		return
	}

	ctx.Response.JSON(eel.Map{
		"id": customer.Id,
	})
}

