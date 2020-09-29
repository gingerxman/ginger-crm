package point

import (
	"context"
	"github.com/gingerxman/eel"
)

type EncodePointRuleService struct {
	eel.ServiceBase
}

func NewEncodePointRuleService(ctx context.Context) *EncodePointRuleService {
	service := new(EncodePointRuleService)
	service.Ctx = ctx
	return service
}

func (this *EncodePointRuleService) Encode(rule *PointRule) *RPointRule {
	rPointRule := &RPointRule{
		Id: rule.Id,
		CorpId: rule.CorpId,
		Type: rule.GetTypeText(),
		Stage: rule.GetStageText(),
		
		Name: rule.Name,
		Detail: rule.Detail,
		Data: rule.Data,
		Point: rule.Point,
		
		IsSystemRule: rule.IsSystemRule,
		IsEnabled: rule.IsEnabled,
		IsDeleted: rule.IsDeleted,
		CreatedAt: rule.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: rule.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	
	return rPointRule
}

func (this *EncodePointRuleService) EncodeMany(rules []*PointRule) []*RPointRule {
	rows := make([]*RPointRule, len(rules))
	
	for i, rule := range rules {
		rows[i] = this.Encode(rule)
	}
	
	return rows
}