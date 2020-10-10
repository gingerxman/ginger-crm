package blog

import (
	"github.com/gingerxman/eel"
	"time"
)


const PR_STAGE_PRECONDITION = 1
const PR_STAGE_OUTPUT = 2
const PR_STAGE_ALL = 3
var PR__STAGE2STR = map[int]string {
	PR_STAGE_PRECONDITION: "precondition",
	PR_STAGE_OUTPUT: "output",
	PR_STAGE_ALL: "all",
}
var PR__STR2STAGE = map[string]int {
	"precondition": PR_STAGE_PRECONDITION,
	"output": PR_STAGE_OUTPUT,
	"all": PR_STAGE_ALL,
}

//PointRule 自定义积分规则pre
const PR_TYPE_TRADE = 1
const PR_TYPE_MONEY = 2
const PR_TYPE_VALID_DAYS = 3
const PR_TYPE_MAX_PER_DAY = 4
const PR_TYPE_PROTECT_DAYS = 5
var PR__TYPE2STR = map[int]string {
	PR_TYPE_TRADE: "trade",
	PR_TYPE_MONEY: "money",
	PR_TYPE_VALID_DAYS: "valid_days",
	PR_TYPE_MAX_PER_DAY: "max_per_day",
	PR_TYPE_PROTECT_DAYS: "protect_days",
}
var PR__STR2TYPE = map[string]int {
	"trade": PR_TYPE_TRADE,
	"money": PR_TYPE_MONEY,
	"valid_days": PR_TYPE_VALID_DAYS,
	"max_per_day": PR_TYPE_MAX_PER_DAY,
	"protect_days": PR_TYPE_PROTECT_DAYS,
}
var PR__TYPE2STAGE = map[string]int {
	"trade": PR_STAGE_ALL,
	"money": PR_STAGE_ALL,
	"valid_days": PR_STAGE_OUTPUT,
	"max_per_day": PR_STAGE_PRECONDITION,
	"protect_days": PR_STAGE_OUTPUT,
}
type PointRule struct {
	eel.Model
	Type int
	Stage int
	CorpId int `gorm:"index"`
	IsSystemRule bool `gorm:"default:false"`
	Name string `gorm:"size:128"`
	Detail string `gorm:"size:512"`
	Data string `gorm:"type:text"` // json data
	Point int // 奖励积分值
	IsForAllProduct bool // 规则是否适用于全部商品
	IsEnabled bool `gorm:"default:true"`
	IsDeleted bool `gorm:"default:false"`
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

//PointLog
type PointLog struct {
	eel.Model
	UserId int `gorm:"index"`
	CorpId int `gorm:"index"`
	OrderBid string `gorm:"index"`
	Point int // 积分数量
	Data string // 业务数据
	InvalidateDate time.Time `gorm:"type:date"`
}
func (self *PointLog) TableName() string {
	return "point_log"
}

//UserPendingTradeCount
type UserPendingTradeCount struct {
	eel.Model
	UserId int `gorm:"index"`
	PendingCount int
}
func (self *UserPendingTradeCount) TableName() string {
	return "point_user_pending_trade_count"
}


func init() {
	eel.RegisterModel(new(PointRule))
	eel.RegisterModel(new(PointRuleHasProduct))
	eel.RegisterModel(new(PointLog))
	eel.RegisterModel(new(UserPendingTradeCount))
}
