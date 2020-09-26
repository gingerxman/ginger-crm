package customer

import (
	"context"
	"errors"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business"
	"github.com/gingerxman/ginger-crm/business/account"
	m_customer "github.com/gingerxman/ginger-crm/models/customer"
)

type consumptionRecord struct {
	ConsumeCount int
	ConsumeMoney int
	LatestConsumeTime string
}

type Customer struct {
	eel.EntityBase
	Id                int
	
	UserId int
	CorpId int
	Unionid string
	Code string
	
	//基本信息
	Name string
	Avatar string
	Sex string
	
	//其他信息
	Source string
	
	// 消费记录
	ConsumptionRecord *consumptionRecord
}

func NewCustomerForCorp(ctx context.Context, userIface business.IUser, corp business.ICorp) (*Customer, error) {
	o := eel.GetOrmFromContext(ctx)
	if o.Model(&m_customer.Customer{}).Where(eel.Map{
		"corp_id": corp.GetId(),
		"user_id": userIface.GetId(),
	}).Exist() {
		var model m_customer.Customer
		db := o.Model(&m_customer.Customer{}).Where(eel.Map{
			"corp_id": corp.GetId(),
			"user_id": userIface.GetId(),
		}).Find(&model)
		
		if db.Error != nil {
			return nil, db.Error
		} else {
			return NewCustomerFromModel(ctx, &model), nil
		}
	} else {
		user := account.NewUserRepository(ctx).GetUserById(userIface.GetId())
		
		if user == nil {
			return nil, errors.New("customer:fetch_user_fail")
		}
		
		model := &m_customer.Customer{
			UserId: user.Id,
			CorpId: corp.GetId(),
			Unionid: user.Unionid,
			Code: user.Code,
			Name: user.Name,
			Avatar: user.Avatar,
			Sex: user.Sex,
			Source: user.Source,
		}
		
		db := o.Create(model)
		if db.Error != nil {
			eel.Logger.Error(db.Error)
			return nil, db.Error
		}
		
		return NewCustomerFromModel(ctx, model), nil
	}
}

func NewCustomerFromModel(ctx context.Context, model *m_customer.Customer) *Customer {
	instance := new(Customer)
	
	instance.Ctx = ctx
	instance.Id = model.Id
	instance.UserId = model.UserId
	instance.CorpId = model.CorpId
	instance.Unionid = model.Unionid
	instance.Code = model.Code
	instance.Name = model.Name
	instance.Avatar = model.Avatar
	instance.Sex = model.Sex
	instance.Source = model.Source
	instance.ConsumptionRecord = &consumptionRecord{}
	
	return instance
}

func init() {
}
