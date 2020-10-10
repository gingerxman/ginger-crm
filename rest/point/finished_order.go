package point

import (
	"fmt"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business/account"
	"github.com/gingerxman/ginger-crm/business/finance"
	b_point "github.com/gingerxman/ginger-crm/business/point"
)

type FinishedOrder struct {
	eel.RestResource
}

func (this *FinishedOrder) Resource() string {
	return "point.finished_order"
}

func (this *FinishedOrder) GetLockKey(ctx *eel.Context) string {
	bCtx := ctx.GetBusinessContext()
	user := account.GetUserFromContext(bCtx)
	
	return fmt.Sprintf("ginger-crm:point.finished_order:%d", user.GetId())
}

func (this *FinishedOrder) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"bid",
			"money:int",
		},
	}
}

func (this *FinishedOrder) Put(ctx *eel.Context) {
	bCtx := ctx.GetBusinessContext()
	
	corp := account.GetCorpFromContext(bCtx)
	if corp.Id == 0 {
		ctx.Response.Error("point_rule:invalid_corp", "corp_id(0)")
		return
	}

	req := ctx.Request
	bid := req.GetString("bid")
	money, _ := req.GetInt("money")
	
	user := account.GetUserFromContext(bCtx)
	point, matchedRules := b_point.NewRuleEngine(bCtx).RunForCorp(corp, bid, money)
	if point > 0 {
		err := finance.NewFinanceService(bCtx).DoTransfer(user, point, bid)
		if err == nil {
			b_point.RecordPointLog(bCtx, corp, user, bid, money, matchedRules, point)
			
			// 如果matchedRules中有trade rule，则清空pending trade count表，否则，递增
			hasTradeRule := false
			for _, rule := range matchedRules {
				if rule.IsTradeRule() {
					hasTradeRule = true
					b_point.NewPointService(bCtx).ResetPendingTradeCountForUser(user)
					break
				}
			}
			if !hasTradeRule {
				// 没有匹配trade rule，记录消费
				b_point.NewPointService(bCtx).RecordTradeForUser(user)
			}
		} else {
			eel.Logger.Error(err)
			point = 0
		}
	} else {
		// 没有匹配规则，记录消费
		b_point.NewPointService(bCtx).RecordTradeForUser(user)
	}
	
	ctx.Response.JSON(eel.Map{
		"point": point,
	})
}
