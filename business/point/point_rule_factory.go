package point

import (
	"context"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business"
	m_point "github.com/gingerxman/ginger-crm/models/point"
)


type PointRuleFactory struct {
	eel.ServiceBase
}

func NewPointRuleFactory(ctx context.Context) *PointRuleFactory {
	service := new(PointRuleFactory)
	service.Ctx = ctx
	return service
}

func (this *PointRuleFactory) InitSystemRulesForCorp(corp business.ICorp) {
	o := eel.GetOrmFromContext(this.Ctx)
	
	// 积分保护期
	model := &m_point.PointRule{
		Type: m_point.PR_TYPE_PROTECT_DAYS,
		Stage: m_point.PR_STAGE_OUTPUT,
		CorpId: corp.GetId(),
		Name: "积分保护期",
		Detail: "",
		Data: "{}",
		Point: 0,
		IsSystemRule: true,
		IsEnabled: false,
		IsDeleted: false,
	}
	
	db := o.Create(model)
	if db.Error != nil {
		eel.Logger.Error(db.Error)
	}
	
	// 积分获取上限
	model = &m_point.PointRule{
		Type: m_point.PR_TYPE_MAX_PER_DAY,
		Stage: m_point.PR_STAGE_PRECONDITION,
		CorpId: corp.GetId(),
		Name: "积分获取上限",
		Detail: "",
		Data: "{}",
		Point: 0,
		IsSystemRule: true,
		IsEnabled: false,
		IsDeleted: false,
	}
	
	db = o.Create(model)
	if db.Error != nil {
		eel.Logger.Error(db.Error)
	}
	
	// 积分有效期
	model = &m_point.PointRule{
		Type: m_point.PR_TYPE_VALID_DAYS,
		Stage: m_point.PR_STAGE_PRECONDITION,
		CorpId: corp.GetId(),
		Name: "积分有效期",
		Detail: "",
		Data: "{}",
		Point: 0,
		IsSystemRule: true,
		IsEnabled: false,
		IsDeleted: false,
	}
	
	db = o.Create(model)
	if db.Error != nil {
		eel.Logger.Error(db.Error)
	}
}

func (this *PointRuleFactory) CreateRuleForCorp(corp business.ICorp, name string, ruleType string, data interface{}, point int) (*PointRule, error){
	o := eel.GetOrmFromContext(this.Ctx)
	
	// 积分有效期
	model := &m_point.PointRule{
		Type: m_point.PR__STR2TYPE[ruleType],
		Stage: m_point.PR__TYPE2STAGE[ruleType],
		CorpId: corp.GetId(),
		Name: name,
		Detail: "",
		Data: eel.ToJsonString(data),
		Point: point,
		IsSystemRule: false,
		IsEnabled: true,
		IsDeleted: false,
	}
	
	db := o.Create(model)
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return nil, db.Error
	}
	
	return NewPointRuleFromModel(this.Ctx, model), nil
}

func init() {
}
