package point

import (
	"fmt"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business/account"
	b_point "github.com/gingerxman/ginger-crm/business/point"
)

type PointRules struct {
	eel.RestResource
}

func (this *PointRules) Resource() string {
	return "point.point_rules"
}

func (this *PointRules) GetParameters() map[string][]string {
	return map[string][]string{
		"GET":  []string{"?filters:json"},
	}
}

func (this *PointRules) Get(ctx *eel.Context) {
	req := ctx.Request
	bCtx := ctx.GetBusinessContext()
	
	filters := req.GetOrmFilters()
	pageInfo := req.GetPageInfo()
	
	fmt.Println(filters)
	fmt.Println(pageInfo)
	
	corp := account.GetCorpFromContext(bCtx)
	rules, nextPageInfo := b_point.NewPointRuleRepository(bCtx).GetPagedPointRulesForCorp(corp, filters, pageInfo)
	
	b_point.NewFillPointRuleService(bCtx).Fill(rules, eel.FillOption{})
	datas := b_point.NewEncodePointRuleService(bCtx).EncodeMany(rules)

	ctx.Response.JSON(eel.Map{
		"rules": datas,
		"pageinfo": nextPageInfo.ToMap(),
	})
}

