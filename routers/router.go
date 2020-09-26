package routers

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/eel/handler/rest/console"
	"github.com/gingerxman/eel/handler/rest/op"
	"github.com/gingerxman/ginger-crm/rest/customer"
	"github.com/gingerxman/ginger-crm/rest/dev"
)

func init() {
	eel.RegisterResource(&console.Console{})
	eel.RegisterResource(&op.Health{})
	
	eel.RegisterResource(&customer.Customer{})
	eel.RegisterResource(&customer.Customers{})
	
	eel.RegisterResource(&dev.BDDReset{})
}