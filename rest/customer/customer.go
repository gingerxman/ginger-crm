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
		"PUT":  []string{},
	}
}

func (this *Customers) Put(ctx *eel.Context) {
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

