package customer

import (
	"context"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business"
	m_customer "github.com/gingerxman/ginger-crm/models/customer"
)

type CustomerRepository struct {
	eel.ServiceBase
}

func NewCustomerRepository(ctx context.Context) *CustomerRepository {
	service := new(CustomerRepository)
	service.Ctx = ctx
	return service
}

func (this *CustomerRepository) GetCustomers(filters eel.Map, orderExprs ...string) []*Customer {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_customer.Customer
	db := o.Model(&m_customer.Customer{})
	if len(filters) > 0 {
		db = db.Where(filters)
	}
	for _, expr := range orderExprs {
		db = db.Order(expr)
	}
	db = db.Find(&models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return make([]*Customer, 0)
	}
	
	instances := make([]*Customer, 0)
	for _, model := range models {
		instances = append(instances, NewCustomerFromModel(this.Ctx, model))
	}
	return instances
}

func (this *CustomerRepository) GetPagedCustomers(filters eel.Map, page *eel.PageInfo, orderExprs ...string) ([]*Customer, eel.INextPageInfo) {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_customer.Customer
	db := o.Model(&m_customer.Customer{})
	if len(filters) > 0 {
		db = db.Where(filters)
	}
	for _, expr := range orderExprs {
		db = db.Order(expr)
	}
	
	paginateResult, db := eel.Paginate(db, page, &models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return nil, paginateResult
	}
	
	instances := make([]*Customer, 0)
	for _, model := range models {
		instances = append(instances, NewCustomerFromModel(this.Ctx, model))
	}
	return instances, paginateResult
}

func (this *CustomerRepository) GetPagedCustomersForCorp(corp business.ICorp, filters eel.Map, page *eel.PageInfo) ([]*Customer, eel.INextPageInfo) {
	filters["corp_id"] = corp.GetId()
	return this.GetPagedCustomers(filters, page, "-id")
}

func (this *CustomerRepository) GetCustomerByIdInCorp(corp business.ICorp, customerId int) *Customer {
	filters := eel.Map{
		"corp_id": corp.GetId(),
		"id": customerId,
	}
	
	customers := this.GetCustomers(filters)
	if len(customers) > 0 {
		return customers[0]
	} else {
		return nil
	}
}

func (this *CustomerRepository) GetCustomerByUserIdInCorp(corp business.ICorp, user business.IUser) *Customer {
	filters := eel.Map{
		"corp_id": corp.GetId(),
		"user_id": user.GetId(),
	}
	
	customers := this.GetCustomers(filters)
	if len(customers) > 0 {
		return customers[0]
	} else {
		return nil
	}
}

func init() {
}
