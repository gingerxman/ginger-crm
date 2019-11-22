package customer

import (
	"encoding/json"
	"fmt"
	"github.com/gingerxman/eel"
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
	
	// corp := account.GetCorpFromContext(bCtx)
	customers := customer.NewCustomerRepository(bCtx).GetByConsumptionRecord()
	
	datas := customer.NewEncodeCustomerService(bCtx).EncodeManyLintCustomers(customers)

	bytes, _ := json.Marshal(datas)
	fmt.Println(string(bytes))
	ctx.Response.JSON(eel.Map{
		"customers": datas,
		"pageinfo": nil, // nextPageInfo.ToMap(),
	})
}

