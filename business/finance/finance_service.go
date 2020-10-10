package finance

import (
	"context"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business"
)

type FinanceService struct {
	eel.RepositoryBase
}

func NewFinanceService(ctx context.Context) *FinanceService {
	repository := new(FinanceService)
	repository.Ctx = ctx
	return repository
}

func (this *FinanceService) DoTransfer(destUser business.IUser, amount int, bid string) error {
	_, err := eel.NewResource(this.Ctx).Put("ginger-finance", "imoney.transfer", eel.Map{
		"bid": bid,
		"amount": amount,
		"source_user_id": 0,
		"dest_user_id": destUser.GetId(),
		"imoney_code": "point",
	})
	
	if err != nil {
		eel.Logger.Error(err)
		return err
	}
	
	return nil
}

func init() {
}
