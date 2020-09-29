package point

import (
	"context"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business"
	m_point "github.com/gingerxman/ginger-crm/models/point"
)

type PointRuleRepository struct {
	eel.ServiceBase
}

func NewPointRuleRepository(ctx context.Context) *PointRuleRepository {
	service := new(PointRuleRepository)
	service.Ctx = ctx
	return service
}

func (this *PointRuleRepository) GetPointRules(filters eel.Map, orderExprs ...string) []*PointRule {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_point.PointRule
	db := o.Model(&m_point.PointRule{})
	filters["is_deleted"] = false
	if len(filters) > 0 {
		db = db.Where(filters)
	}
	for _, expr := range orderExprs {
		db = db.Order(expr)
	}
	db = db.Find(&models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return make([]*PointRule, 0)
	}
	
	instances := make([]*PointRule, 0)
	for _, model := range models {
		instances = append(instances, NewPointRuleFromModel(this.Ctx, model))
	}
	return instances
}

func (this *PointRuleRepository) GetPagedPointRules(filters eel.Map, page *eel.PageInfo, orderExprs ...string) ([]*PointRule, eel.INextPageInfo) {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_point.PointRule
	db := o.Model(&m_point.PointRule{})
	filters["is_deleted"] = false
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
	
	instances := make([]*PointRule, 0)
	for _, model := range models {
		instances = append(instances, NewPointRuleFromModel(this.Ctx, model))
	}
	return instances, paginateResult
}

func (this *PointRuleRepository) GetPagedPointRulesForCorp(corp business.ICorp, filters eel.Map, page *eel.PageInfo) ([]*PointRule, eel.INextPageInfo) {
	filters["corp_id"] = corp.GetId()
	rules, nextPageInfo := this.GetPagedPointRules(filters, page, "-id")
	
	return rules, nextPageInfo
}

func (this *PointRuleRepository) GetPointRuleInCorp(corp business.ICorp, ruleId int) *PointRule {
	filters := eel.Map{
		"id": ruleId,
		"corp_id": corp.GetId(),
	}
	rules := this.GetPointRules(filters)
	
	if len(rules) > 0 {
		return rules[0]
	} else {
		return nil
	}
}


func init() {
}
