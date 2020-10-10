package point

import (
	"context"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business"
	m_point "github.com/gingerxman/ginger-crm/models/point"
)

type PointLogRepository struct {
	eel.ServiceBase
}

func NewPointLogRepository(ctx context.Context) *PointLogRepository {
	service := new(PointLogRepository)
	service.Ctx = ctx
	return service
}

func (this *PointLogRepository) GetPointLogs(filters eel.Map, orderExprs ...string) []*PointLog {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_point.PointLog
	db := o.Model(&m_point.PointLog{})
	if len(filters) > 0 {
		db = db.Where(filters)
	}
	for _, expr := range orderExprs {
		db = db.Order(expr)
	}
	db = db.Find(&models)
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return make([]*PointLog, 0)
	}
	
	instances := make([]*PointLog, 0)
	for _, model := range models {
		instances = append(instances, NewPointLogFromModel(this.Ctx, model))
	}
	return instances
}

func (this *PointLogRepository) GetPagedPointLogs(filters eel.Map, page *eel.PageInfo, orderExprs ...string) ([]*PointLog, eel.INextPageInfo) {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_point.PointLog
	db := o.Model(&m_point.PointLog{})
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
	
	instances := make([]*PointLog, 0)
	for _, model := range models {
		instances = append(instances, NewPointLogFromModel(this.Ctx, model))
	}
	return instances, paginateResult
}

func (this *PointLogRepository) GetPagedPointLogsForCorp(corp business.ICorp, filters eel.Map, page *eel.PageInfo) ([]*PointLog, eel.INextPageInfo) {
	filters["corp_id"] = corp.GetId()
	logs, nextPageInfo := this.GetPagedPointLogs(filters, page, "-id")
	
	return logs, nextPageInfo
}

func (this *PointLogRepository) GetPagedPointLogsForUserInCorp(corp business.ICorp, user business.IUser, filters eel.Map, page *eel.PageInfo) ([]*PointLog, eel.INextPageInfo) {
	filters["corp_id"] = corp.GetId()
	filters["user_id"] = user.GetId()
	logs, nextPageInfo := this.GetPagedPointLogs(filters, page, "-id")
	
	return logs, nextPageInfo
}

func (this *PointLogRepository) GetPointLogsForUserInCorp(corp business.ICorp, user business.IUser, filters eel.Map) []*PointLog {
	filters["corp_id"] = corp.GetId()
	filters["user_id"] = user.GetId()
	logs := this.GetPointLogs(filters, "-id")
	
	return logs
}


func init() {
}
