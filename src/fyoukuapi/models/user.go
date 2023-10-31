package models

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type User struct {
	Id       int
	Name     string
	Password string
	AddTime  int64
	Status   int
	Mobile   string
	Avatar   string
}

func init() {
	orm.RegisterModel(new(User))
}

// 保存用户信息
func UserSave(mobile string, password string) (User, error) {
	query := orm.NewOrm()
	user := User{Name: "", Password: password, Mobile: mobile, Status: 1, AddTime: time.Now().Unix()}
	err := query.Read(&user, "Mobile")
	if err == nil {
		return User{}, errors.New("手机号已被使用")
	}
	_, err = query.Insert(&user)
	if err != nil {
		return User{}, errors.New("内部异常")
	}
	return User{}, nil
}

// 获取用户信息
func GetUserinfo(condition map[string]interface{}) (User, bool) {
	// 创建查询条件
	query := orm.NewOrm().QueryTable("user")

	if condition["id"] != nil {
		query = query.Filter("Id", condition["id"])
	}
	if condition["name"] != nil {
		query = query.Filter("Name", condition["name"])
	}
	if condition["mobile"] != nil {
		query = query.Filter("Mobile", condition["mobile"])
	}
	if condition["password"] != nil {
		query = query.Filter("Password", condition["password"])
	}

	user := User{}
	err := query.One(&user)
	if err != nil {
		return user, false
	}
	return user, true
}
