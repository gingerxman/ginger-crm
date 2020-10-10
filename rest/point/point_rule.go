package point

import (
	"fmt"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business/account"
	b_point "github.com/gingerxman/ginger-crm/business/point"
)

type PointRule struct {
	eel.RestResource
}

func (this *PointRule) Resource() string {
	return "point.point_rule"
}

func (this *PointRule) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"id:int"},
		"PUT": []string{
			"type",
			"?name",
			"data:json",
			"point:int",
			"?is_system_rule:bool",
		},
		"POST": []string{
			"id:int",
			"type",
			"data:json",
			"point:int",
		},
		"DELETE": []string{
			"id:int",
		},
	}
}

func (this *PointRule) Get(ctx *eel.Context) {
	req := ctx.Request
	bCtx := ctx.GetBusinessContext()
	
	corp := account.GetCorpFromContext(bCtx)
	ruleId, _ := req.GetInt("id")
	rule := b_point.NewPointRuleRepository(bCtx).GetPointRuleInCorp(corp, ruleId)
	if rule == nil {
		ctx.Response.Error("point_rule:invalid_rule", fmt.Sprintf("id(%d)", ruleId))
		return
	}
	
	b_point.NewFillPointRuleService(bCtx).FillOne(rule, eel.FillOption{})
	data := b_point.NewEncodePointRuleService(bCtx).Encode(rule)
	
	ctx.Response.JSON(data)
}

func (this *PointRule) Put(ctx *eel.Context) {
	bCtx := ctx.GetBusinessContext()
	
	corp := account.GetCorpFromContext(bCtx)
	if corp.Id == 0 {
		ctx.Response.Error("point_rule:invalid_corp", "corp_id(0)")
		return
	}

	req := ctx.Request
	name := req.GetString("name", "custom")
	ruleType := req.GetString("type")
	data := req.GetJSON("data")
	point, _ := req.GetInt("point", 0)
	rule, err := b_point.NewPointRuleFactory(bCtx).CreateRuleForCorp(corp, name, ruleType, data, point)
	if err != nil {
		ctx.Response.Error("point_rule:create_fail", err.Error())
		return
	}

	ctx.Response.JSON(eel.Map{
		"id": rule.Id,
	})
}

func (this *PointRule) Post(ctx *eel.Context) {
	bCtx := ctx.GetBusinessContext()
	
	corp := account.GetCorpFromContext(bCtx)
	if corp.Id == 0 {
		ctx.Response.Error("point_rule:invalid_corp", "corp_id(0)")
		return
	}
	
	req := ctx.Request
	ruleId, _ := req.GetInt("id")
	rule := b_point.NewPointRuleRepository(bCtx).GetPointRuleInCorp(corp, ruleId)
	if rule == nil {
		ctx.Response.Error("point_rule:invalid_rule", fmt.Sprintf("id(%d)", ruleId))
		return
	}
	
	
	ruleType := req.GetString("type")
	data := req.GetJSON("data")
	point, _ := req.GetInt("point", 0)
	rule.Update(ruleType, data, point)
	
	ctx.Response.JSON(eel.Map{})
}

func (this *PointRule) Delete(ctx *eel.Context) {
	bCtx := ctx.GetBusinessContext()
	
	corp := account.GetCorpFromContext(bCtx)
	if corp.Id == 0 {
		ctx.Response.Error("point_rule:invalid_corp", "corp_id(0)")
		return
	}
	
	req := ctx.Request
	ruleId, _ := req.GetInt("id")
	rule := b_point.NewPointRuleRepository(bCtx).GetPointRuleInCorp(corp, ruleId)
	if rule == nil {
		ctx.Response.Error("point_rule:invalid_rule", fmt.Sprintf("id(%d)", ruleId))
		return
	}
	
	err := rule.Delete()
	if err != nil {
		eel.Logger.Error(err)
		ctx.Response.Error("point_rule:delete_fail", err.Error())
		return
	}
	
	ctx.Response.JSON(eel.Map{})
}
