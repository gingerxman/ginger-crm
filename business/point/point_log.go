package point

import (
	"context"
	"encoding/json"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business"
	"github.com/gingerxman/ginger-crm/business/account"
	m_point "github.com/gingerxman/ginger-crm/models/point"
	"time"
)


type PointLog struct {
	eel.EntityBase
	Id int
	
	Data map[string]interface{} // 业务数据
	Point int // 积分数量
	
	CreatedAt  time.Time
	
	UserId int
	User *account.User
	
	CorpId int
}

func GenerateLogForRule(corp business.ICorp, user business.IUser, orderBid string, rule *PointRule, point int) *m_point.PointLog {
	if point < rule.Point {
		rule.Data["limited_by_max_per_day_rule"] = true
	}
	
	model := &m_point.PointLog{
		UserId: user.GetId(),
		CorpId: corp.GetId(),
		OrderBid: orderBid,
		Point: rule.Point,
		Data: eel.ToJsonString(rule.Data),
		InvalidateDate: eel.ParseTime("2000-01-01 00:00:00"),
	}
	
	if rule.IsMaxPerDayRule() {
		now := time.Now()
		days := int(rule.Data["count"].(float64))
		invalidateDate := now.Add(time.Hour * 24 * time.Duration(days))
		model.InvalidateDate = invalidateDate
	}
	
	return model
}

// RecordPointLog : 记录一次积分日志
func RecordPointLog(ctx context.Context, corp business.ICorp, user business.IUser, orderBid string, orderMoney int, rules []*PointRule, point int) {
	o := eel.GetOrmFromContext(ctx)
	
	ruleDatas := make([]interface{}, 0)
	for _, rule := range rules {
		data := eel.Map{}
		data["type"] = rule.GetTypeText()
		data["data"] = rule.Data
		data["point"] = rule.Point
		ruleDatas = append(ruleDatas, data)
	}
	
	if !o.Model(&m_point.PointLog{}).Where("order_bid", orderBid).Exist() {
		model := &m_point.PointLog{
			UserId: user.GetId(),
			CorpId: corp.GetId(),
			OrderBid: orderBid,
			Point: point,
			Data: eel.ToJsonString(eel.Map{
				"order_money": orderMoney,
				"rules": ruleDatas,
			}),
		}
		
		db := o.Create(model)
		if db.Error != nil {
			eel.Logger.Error(db.Error)
		}
	}
}

func NewPointLogFromModel(ctx context.Context, model *m_point.PointLog) *PointLog {
	instance := new(PointLog)
	
	instance.Ctx = ctx
	instance.Id = model.Id
	instance.UserId = model.UserId
	instance.CorpId = model.CorpId
	instance.Point = model.Point
	instance.CreatedAt = model.CreatedAt

	err := json.Unmarshal([]byte(model.Data), &instance.Data)
	if err != nil {
		eel.Logger.Error(err)
	}
	
	return instance
}

func init() {
}
