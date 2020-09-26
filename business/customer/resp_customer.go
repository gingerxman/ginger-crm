package customer


type RCustomer struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Unionid string `json:"unionid"`
	Code string `json:"code"`
	
	// 客户信息
	Sex string `json:"sex"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	Source string `json:"source"`

	// 消费记录
	ConsumeCount int `json:"consume_count"`
	ConsumeMoney int `json:"consume_money"`
	LatestConsumeTime string `json:"latest_consume_time"`
}
