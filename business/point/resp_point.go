package point


type RPointRule struct {
	Id int `json:"id"`
	CorpId int `json:"corp_id"`
	Type string `json:"type"`
	Stage string `json:"stage"`
	
	IsSystemRule bool `json:"is_system_rule"`
	Name string `json:"name"`
	Detail string `json:"detail"`
	Data interface{} `json:"data"`
	Point int `json:"point"`
	
	IsEnabled bool `json:"is_enabled"`
	IsDeleted bool `json:"is_deleted"`
	
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type RPointLog struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	CorpId int `json:"corp_id"`
	Point int `json:"point"'`
	
	User interface{} `json:"user"`
	Data interface{} `json:"data"`
	
	CreatedAt string `json:"created_at"`
}
