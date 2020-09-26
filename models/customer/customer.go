package blog

import (
	"github.com/gingerxman/eel"
)

// Customer
type Customer struct {
	eel.Model
	UserId int `gorm:"index"`
	CorpId int `gorm:"index"`
	Unionid string `gorm:"size:255"`
	Code string `gorm:"size:32"`
	
	//基本信息
	Name string `gorm:"size:52"`
	Avatar string `gorm:"size:1024"`
	Sex string
	
	//其他信息
	Source string `gorm:"size:52"`
}
func (this *Customer) TableName() string {
	return "customer_customer"
}

func init() {
	eel.RegisterModel(new(Customer))
}
