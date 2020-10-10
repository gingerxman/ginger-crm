package point

import (
	"context"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business/account"
)

type EncodePointLogService struct {
	eel.ServiceBase
}

func NewEncodePointLogService(ctx context.Context) *EncodePointLogService {
	service := new(EncodePointLogService)
	service.Ctx = ctx
	return service
}

func (this *EncodePointLogService) Encode(log *PointLog) *RPointLog {
	var rUser *account.RUser
	if log.User != nil {
		user := log.User
		rUser = &account.RUser{
			Id: user.Id,
			Name: user.Name,
			Avatar: user.Avatar,
			Sex: user.Sex,
			Code: user.Code,
		}
	}
	
	rPointLog := &RPointLog{
		Id: log.Id,
		CorpId: log.CorpId,
		Point: log.Point,
		
		User: rUser,
		Data: log.Data,
		
		CreatedAt: log.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	
	return rPointLog
}

func (this *EncodePointLogService) EncodeMany(logs []*PointLog) []*RPointLog {
	rows := make([]*RPointLog, len(logs))
	
	for i, log := range logs {
		rows[i] = this.Encode(log)
	}
	
	return rows
}