package account

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gingerxman/eel"
)

type UserRepository struct {
	eel.ServiceBase
}

func NewUserRepository(ctx context.Context) *UserRepository {
	service := new(UserRepository)
	service.Ctx = ctx
	return service
}

func (this *UserRepository) makeUsers(userDatas []interface{}) []*User {
	users := make([]*User, 0)
	for _, userData := range userDatas {
		userJson := userData.(map[string]interface{})
		id, _ := userJson["id"].(json.Number).Int64()
		user := NewUserFromOnlyId(this.Ctx, int(id))
		user.Unionid = userJson["unionid"].(string)
		user.Name = userJson["name"].(string)
		user.Avatar = userJson["avatar"].(string)
		user.Sex = userJson["sex"].(string)
		user.Code = userJson["code"].(string)
		user.Source = userJson["source"].(string)
		
		users = append(users, user)
	}
	
	return users
}

func (this *UserRepository) GetUsers(ids []int) []*User {
	options := make(map[string]interface{})
	options["with_role_info"] = true
	resp, err := eel.NewResource(this.Ctx).Get("ginger-account", "user.users", eel.Map{
		"ids": eel.ToJsonString(ids),
	})

	if err != nil {
		eel.Logger.Error(err)
		return nil
	}

	respData := resp.Data()
	userDatas := respData.Get("users")
	fmt.Println(userDatas)
	return this.makeUsers(userDatas.MustArray())
}

func (this *UserRepository) GetUserById(id int) *User {
	users := this.GetUsers([]int{id})
	
	if len(users) > 0 {
		return users[0]
	} else {
		return nil
	}
}

func init() {
}
