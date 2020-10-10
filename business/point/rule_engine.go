package point

import (
	"context"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business"
	"github.com/gingerxman/ginger-crm/business/account"
	m_point "github.com/gingerxman/ginger-crm/models/point"
)

type Order struct {
	Bid string
	Money int
}

type RuleEngine struct {
	eel.ServiceBase
}

func NewRuleEngine(ctx context.Context) *RuleEngine {
	service := new(RuleEngine)
	service.Ctx = ctx
	return service
}

func (this *RuleEngine) RunForCorp(corp business.ICorp, bid string, money int) (int, []*PointRule) {
	rules := NewPointRuleRepository(this.Ctx).GetPointRules(eel.Map{
		"corp_id": corp.GetId(),
	})
	
	user := account.GetUserFromContext(this.Ctx)
	usedRules := make([]*PointRule, 0)
	
	// 获得当天剩余还可积累的money rule余额
	// 最后处理"积分上限"规则
	remainMoneyRulePoint := 99999999999
	for _, rule := range rules {
		if rule.Type == m_point.PR_TYPE_MAX_PER_DAY && rule.IsActive() {
			maxPoint := int(rule.Data["count"].(float64)) * 100
			todayPoint := NewPointService(this.Ctx).GetTodayPointForUserInCorp(corp, user)
			remainMoneyRulePoint = maxPoint - todayPoint
			usedRules = append(usedRules, rule)
		}
	}
	
	// 匹配规则
	moneyRules := make([]*PointRule, 0)
	tradeRules := make([]*PointRule, 0)
	for _, rule := range rules {
		if rule.Type == m_point.PR_TYPE_MONEY {
			baseMoney64, _ := rule.Data["count"].(float64)
			baseMoney := int(baseMoney64)
			
			if money > baseMoney {
				moneyRules = append(moneyRules, rule)
			}
		} else if rule.Type == m_point.PR_TYPE_TRADE {
			targetTradeCount64, _ := rule.Data["count"].(float64)
			targetTradeCount := int(targetTradeCount64)
			
			pointService := NewPointService(this.Ctx)
			pendingTradeCount := pointService.GetPendingTradingCountForUser(user)
			
			if pendingTradeCount + 1 >= targetTradeCount {
				tradeRules = append(tradeRules, rule)
			}
		}
	}
	
	// 寻找门槛最高的money rule
	moneyRulePoint := 0
	if len(moneyRules) > 0 {
		maxMoney := 0
		var maxMoneyRule *PointRule
		for _, moneyRule := range moneyRules {
			money := int(moneyRule.Data["count"].(float64))
			if money > maxMoney {
				maxMoney = money
				maxMoneyRule = moneyRule
			}
		}
		factor := money / maxMoney
		moneyRulePoint = (maxMoneyRule.Point * factor) * 100
		// 根据积分上限，削减money rule point
		if moneyRulePoint > remainMoneyRulePoint {
			moneyRulePoint = remainMoneyRulePoint
			maxMoneyRule.Data["is_limited_by_max_per_day_rule"] = true
		}
		usedRules = append(usedRules, maxMoneyRule)
	}
	
	// 寻找门槛最低的trade rule
	tradeRulePoint := 0
	if len(tradeRules) > 0 {
		minTradeCount := 9999999
		var minTradeRule *PointRule
		for _, tradeRule := range tradeRules {
			tradeCount := int(tradeRule.Data["count"].(float64))
			if tradeCount < minTradeCount {
				minTradeCount = tradeCount
				minTradeRule = tradeRule
			}
		}
		tradeRulePoint = minTradeRule.Point * 100
		usedRules = append(usedRules, minTradeRule)
	}
	
	return tradeRulePoint+moneyRulePoint, usedRules
}

func init() {
}
