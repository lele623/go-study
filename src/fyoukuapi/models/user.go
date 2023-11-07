package models

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
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

// 保存用户信息
func UserSave(mobile string, password string) ([]User, error) {
	var user []User

	num, _ := orm.NewOrm().Raw("select * from user where mobile = ? limit 1", mobile).QueryRows(&user)
	if num != 0 {
		return user, errors.New("手机号已被使用")
	}

	_, err := orm.NewOrm().Raw("insert into user (mobile,password) values (?,?)", mobile, password).Exec()
	if err != nil {
		logs.Error(err)
		return user, errors.New("内部异常")
	}
	return user, nil
}

// 获取用户信息
func GetUserinfo(param map[string]interface{}) ([]User, int64) {
	var (
		user []User
		args []interface{}
	)

	sql := "SELECT * FROM user WHERE 1 = 1"

	if param["mobile"] != "" {
		sql += " AND mobile = ?"
		args = append(args, param["mobile"])
	}
	if param["password"] != "" {
		sql += " AND password = ?"
		args = append(args, param["password"])
	}

	num, _ := orm.NewOrm().Raw(sql, args...).QueryRows(&user)
	return user, num
}
