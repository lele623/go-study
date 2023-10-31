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

func GetUserInfoByMobile(mobile string) User {
	o := orm.NewOrm()
	user := User{Mobile: mobile}
	err := o.Read(&user, "Mobile")
	if err != nil {
		return User{}
	}
	return user
}

func UserSave(mobile string, password string) (User, error) {
	o := orm.NewOrm()
	user := User{Name: "", Password: password, Mobile: mobile, Status: 1, AddTime: time.Now().Unix()}
	err := o.Read(&user, "Mobile")
	if err == nil {
		return User{}, errors.New("手机号已被使用")
	}
	_, err = o.Insert(&user)
	if err != nil {
		return User{}, errors.New("内部异常，注册失败")
	}
	return User{}, nil
}
