package routers

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/eel/handler/rest/console"
	"github.com/gingerxman/eel/handler/rest/op"
	"github.com/gingerxman/ginger-crm/rest/customer"
	"github.com/gingerxman/ginger-crm/rest/dev"
	"github.com/gingerxman/ginger-crm/rest/point"
	"github.com/gingerxman/ginger-crm/rest/system"
)

func init() {
	eel.RegisterResource(&console.Console{})
	eel.RegisterResource(&op.Health{})
	
	// customer
	eel.RegisterResource(&customer.Customer{})
	eel.RegisterResource(&customer.Customers{})
	
	// point
	eel.RegisterResource(&point.PointRule{})
	eel.RegisterResource(&point.PointRules{})
	eel.RegisterResource(&point.FinishedOrder{})
	
	// system
	eel.RegisterResource(&system.SystemInit{})
	
	eel.RegisterResource(&dev.BDDReset{})
}