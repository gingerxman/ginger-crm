package point

import (
	"context"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business/account"
)

type FillPointLogService struct {
	eel.ServiceBase
}

func NewFillPointLogService(ctx context.Context) *FillPointLogService {
	service := new(FillPointLogService)
	service.Ctx = ctx
	return service
}

func (this *FillPointLogService) FillOne(log *PointLog, option eel.FillOption) {
	this.Fill([]*PointLog{log}, option)
}

func (this *FillPointLogService) Fill(logs []*PointLog, option eel.FillOption)  {
	if len(logs) == 0 {
		return
	}
	
	ids := make([]int, 0)
	for _, log := range logs {
		ids = append(ids, log.Id)
	}
	
	this.fillUsers(logs, ids)
}

func (this *FillPointLogService) fillUsers(logs []*PointLog, ids []int) {
	if len(ids) == 0 {
		return
	}
	
	userIds := make([]int, 0)
	
	for _, log := range logs {
		userIds = append(userIds, log.UserId)
	}
	
	users := account.NewUserRepository(this.Ctx).GetUsers(userIds)
	
	//构建<id, user>
	id2user := make(map[int]*account.User)
	for _, user := range users {
		id2user[user.Id] = user
	}
	
	//填充log.User
	for _, log := range logs {
		if user, ok := id2user[log.UserId]; ok {
			log.User =  user
		}
	}
}