package blog

import (
	"github.com/gingerxman/eel"
)

// SystemPointRule 系统的通用积分规则
type SystemPointRule struct {
	eel.Model
	Name string `gorm:"size:128"`
	Detail string `gorm:size:512`
	Data string `gorm:"type:text"`
	IsEnable bool `gorm:"default:true"`
}
func (self *SystemPointRule) TableName() string {
	return "point_system_rule"
}


//PointRule 自定义积分规则
const PR_TYPE_TRADE = 1
const PR_TEYP_MONEY = 2
var PR__TYPE2STR = map[int]string {
	PR_TYPE_TRADE: "trade",
	PR_TEYP_MONEY: "money",
}
var PR__STR2TYPE = map[string]int {
	"trade": PR_TYPE_TRADE,
	"money": PR_TEYP_MONEY,
}
type PointRule struct {
	eel.Model
	Point int // 奖励积分值
	Type int
	TradeCount int
	Money int
	
}
func (self *PointRule) TableName() string {
	return "point_rule"
}


//PointRuleHasProduct
type PointRuleHasProduct struct {
	eel.Model
	RuleId int
	ProductId int // pool product id
}
func (self *PointRuleHasProduct) TableName() string {
	return "point_rule_has_product"
}


func init() {
	eel.RegisterModel(new(SystemPointRule))
	eel.RegisterModel(new(PointRule))
	eel.RegisterModel(new(PointRuleHasProduct))
}
