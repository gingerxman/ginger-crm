package point

import (
	"context"
	"github.com/gingerxman/eel"
)

type FillPointRuleService struct {
	eel.ServiceBase
}

func NewFillPointRuleService(ctx context.Context) *FillPointRuleService {
	service := new(FillPointRuleService)
	service.Ctx = ctx
	return service
}

func (this *FillPointRuleService) FillOne(rule *PointRule, option eel.FillOption) {
	this.Fill([]*PointRule{rule}, option)
}

func (this *FillPointRuleService) Fill(rules []*PointRule, option eel.FillOption)  {
	if len(rules) == 0 {
		return
	}
	
	ids := make([]int, 0)
	for _, rule := range rules {
		ids = append(ids, rule.Id)
	}
}