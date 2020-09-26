package customer

import (
	"fmt"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business/account"
	"github.com/gingerxman/ginger-crm/business/customer"
)

type Customers struct {
	eel.RestResource
}

func (this *Customers) Resource() string {
	return "customer.customers"
}

func (this *Customers) GetParameters() map[string][]string {
	return map[string][]string{
		"GET":  []string{"?filters:json"},
	}
}

func (this *Customers) Get(ctx *eel.Context) {
	req := ctx.Request
	bCtx := ctx.GetBusinessContext()
	
	filters := req.GetOrmFilters()
	pageInfo := req.GetPageInfo()
	
	fmt.Println(filters)
	fmt.Println(pageInfo)
	
	corp := account.GetCorpFromContext(bCtx)
	customers, nextPageInfo := customer.NewCustomerRepository(bCtx).GetPagedCustomersForCorp(corp, filters, pageInfo)
	
	customer.NewFillCustomerService(bCtx).Fill(customers, eel.FillOption{
		"with_consumption_record": true,
	})
	datas := customer.NewEncodeCustomerService(bCtx).EncodeMany(customers)

	ctx.Response.JSON(eel.Map{
		"customers": datas,
		"pageinfo": nextPageInfo.ToMap(),
	})
}

