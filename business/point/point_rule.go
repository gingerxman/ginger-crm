package point

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-crm/business"
	m_point "github.com/gingerxman/ginger-crm/models/point"
	"github.com/gingerxman/gorm"
	"time"
)

type consumptionRecord struct {
	ConsumeCount int
	ConsumeMoney int
	LatestConsumeTime string
}

type PointRule struct {
	eel.EntityBase
	Id int
	
	Type int
	Stage int
	CorpId int
	Name string
	Detail string
	Data map[string]interface{} // 规则业务数据
	Point int // 奖励积分值
	
	IsForAllProduct bool
	IsSystemRule bool
	IsEnabled  bool
	IsDeleted  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (this *PointRule) GetTypeText() string {
	return m_point.PR__TYPE2STR[this.Type]
}

func (this *PointRule) GetStageText() string {
	return m_point.PR__STAGE2STR[this.Stage]
}

func (this *PointRule) Update(ruleType string, data interface{}, point int) {
	o := eel.GetOrmFromContext(this.Ctx)
	db := o.Model(&m_point.PointRule{}).Where("id", this.Id).Update(gorm.Params{
		"type": m_point.PR__STR2TYPE[ruleType],
		"stage": m_point.PR__TYPE2STAGE[ruleType],
		"data": eel.ToJsonString(data),
		"point": point,
		"updated_at": time.Now(),
	})
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
	}
}

func (this *PointRule) Delete() error {
	if this.IsSystemRule {
		return errors.New("不能删除系统规则")
	}
	
	o := eel.GetOrmFromContext(this.Ctx)
	db := o.Where("id", this.Id).Delete(&m_point.PointRule{})
	//db := o.Model(&m_point.PointRule{}).Where("id", this.Id).Update(gorm.Params{
	//	"data": eel.ToJsonString(data),
	//	"point": point,
	//	"updated_at": time.Now(),
	//})
	
	if db.Error != nil {
		eel.Logger.Error(db.Error)
		return db.Error
	}
	
	return nil
}

func (this *PointRule) IsPreconditionRule() bool {
	if this.Stage == m_point.PR_STAGE_PRECONDITION || this.Stage == m_point.PR_STAGE_ALL {
		return true
	}
	
	return false
}

func (this *PointRule) IsOutputRule() bool {
	if this.Stage == m_point.PR_STAGE_OUTPUT || this.Stage == m_point.PR_STAGE_ALL {
		return true
	}
	
	return false
}

func (this *PointRule) IsMaxPerDayRule() bool {
	return this.Type == m_point.PR_TYPE_MAX_PER_DAY
}

func (this *PointRule) IsTradeRule() bool {
	return this.Type == m_point.PR_TYPE_TRADE
}

func (this *PointRule) IsActive() bool {
	if (this.Type == m_point.PR_TYPE_MONEY || this.Type == m_point.PR_TYPE_TRADE) {
		return true
	} else {
		// 系统规则，都检查data中是否有count存在
		data, ok := this.Data["count"]
		if !ok {
			return false
		}
		
		count := int(data.(float64))
		if count == 0 {
			return false
		} else {
			return true
		}
	}
}

func InitSystemRulesForCorp(ctx context.Context, corp business.ICorp) {
	o := eel.GetOrmFromContext(ctx)
	
	// 积分保护期
	model := &m_point.PointRule{
		Type: m_point.PR_TYPE_PROTECT_DAYS,
		Stage: m_point.PR_STAGE_OUTPUT,
		CorpId: corp.GetId(),
		Name: "积分保护期",
		Detail: "",
		Data: "{}",
		Point: 0,
		IsSystemRule: true,
		IsEnabled: false,
		IsDeleted: false,
	}
	
	db := o.Create(model)
	if db.Error != nil {
		eel.Logger.Error(db.Error)
	}
	
	// 积分获取上限
	model = &m_point.PointRule{
		Type: m_point.PR_TYPE_MAX_PER_DAY,
		Stage: m_point.PR_STAGE_PRECONDITION,
		CorpId: corp.GetId(),
		Name: "积分获取上限",
		Detail: "",
		Data: "{}",
		Point: 0,
		IsSystemRule: true,
		IsEnabled: false,
		IsDeleted: false,
	}
	
	db = o.Create(model)
	if db.Error != nil {
		eel.Logger.Error(db.Error)
	}
	
	// 积分有效期
	model = &m_point.PointRule{
		Type: m_point.PR_TYPE_VALID_DAYS,
		Stage: m_point.PR_STAGE_PRECONDITION,
		CorpId: corp.GetId(),
		Name: "积分有效期",
		Detail: "",
		Data: "{}",
		Point: 0,
		IsSystemRule: true,
		IsEnabled: false,
		IsDeleted: false,
	}
	
	db = o.Create(model)
	if db.Error != nil {
		eel.Logger.Error(db.Error)
	}
}

func NewPointRuleFromModel(ctx context.Context, model *m_point.PointRule) *PointRule {
	instance := new(PointRule)
	
	instance.Ctx = ctx
	instance.Id = model.Id
	instance.CorpId = model.CorpId
	instance.Type = model.Type
	instance.Stage = model.Stage
	
	instance.Name = model.Name
	instance.Detail = model.Detail
	instance.Point = model.Point
	
	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(model.Data), &data)
	if err != nil {
		eel.Logger.Error(err)
	}
	instance.Data = data
	
	instance.IsForAllProduct = model.IsForAllProduct
	instance.IsSystemRule = model.IsSystemRule
	instance.IsEnabled = model.IsEnabled
	instance.IsDeleted = model.IsDeleted
	instance.CreatedAt = model.CreatedAt
	instance.UpdatedAt = model.UpdatedAt
	
	return instance
}

func init() {
}
