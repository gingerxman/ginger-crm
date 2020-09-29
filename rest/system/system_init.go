package system

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business/account"
	b_point "github.com/gingerxman/ginger-crm/business/point"
)

type SystemInit struct {
	eel.RestResource
}

func (this *SystemInit) Resource() string {
	return "system.system_init"
}

func (this *SystemInit) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{},
	}
}

func (this *SystemInit) Put(ctx *eel.Context) {
	bCtx := ctx.GetBusinessContext()
	
	corp := account.GetCorpFromContext(bCtx)
	if corp.Id == 0 {
		ctx.Response.Error("point_rule:invalid_corp", "corp_id(0)")
		return
	}
	
	b_point.NewPointRuleFactory(bCtx).InitSystemRulesForCorp(corp)

	ctx.Response.JSON(eel.Map{})
}


