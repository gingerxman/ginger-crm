package point

import (
	"context"
	"fmt"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business"
	m_point "github.com/gingerxman/ginger-crm/models/point"
	"github.com/gingerxman/gorm"
	"time"
)


type PointService struct {
	eel.ServiceBase
}

func NewPointService(ctx context.Context) *PointService {
	service := new(PointService)
	service.Ctx = ctx
	return service
}

// GetPendingTradingCountForUser : 获取user的未处理消费记录
func (this *PointService) GetPendingTradingCountForUser(user business.IUser) int {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var model m_point.UserPendingTradeCount
	if o.Model(&m_point.UserPendingTradeCount{}).Where("user_id", user.GetId()).Exist() {
		db := o.Model(&m_point.UserPendingTradeCount{}).Where("user_id", user.GetId()).Take(&model)
		if db.Error != nil {
			eel.Logger.Error(db.Error)
			return 0
		}
	} else {
		model := &m_point.UserPendingTradeCount{
			UserId: user.GetId(),
			PendingCount: 0,
		}
		
		db := o.Create(model)
		if db.Error != nil {
			eel.Logger.Error(db.Error)
			return 0
		}
	}
	
	return model.PendingCount
}

// RecordTradeForUser : 记录user的一次消费
func (this *PointService) RecordTradeForUser(user business.IUser) error {
	o := eel.GetOrmFromContext(this.Ctx)
	
	db := o.Model(&m_point.UserPendingTradeCount{}).Where("user_id", user.GetId()).Update(gorm.Params{
		"pending_count": gorm.Expr("pending_count + ?", 1),
	})
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return db.Error
	}
	
	return nil
}

// ResetPendingTradeCountForUser : 记录user的未处理消费记录
func (this *PointService) ResetPendingTradeCountForUser(user business.IUser) error {
	o := eel.GetOrmFromContext(this.Ctx)
	
	db := o.Model(&m_point.UserPendingTradeCount{}).Where("user_id", user.GetId()).Update(gorm.Params{
		"pending_count": 0,
	})
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return db.Error
	}
	
	return nil
}

// GetTodayPointForUserInCorp : 获得当天user在corp中领取的积分累积值
func (this *PointService) GetTodayPointForUserInCorp(corp business.ICorp, user business.IUser) int {
	now := time.Now()
	today := now.Format("2006-01-02")
	start_time := fmt.Sprintf("%s 00:00:00", today)
	end_time := fmt.Sprintf("%s 23:59:59", today)
	filters := eel.Map{
		"created_at__gte": start_time,
		"created_at__lte": end_time,
	}
	logs := NewPointLogRepository(this.Ctx).GetPointLogsForUserInCorp(corp, user, filters)
	
	totalPoint := 0
	for _, log := range logs {
		totalPoint += log.Point
	}
	
	return totalPoint
}

func init() {
}
