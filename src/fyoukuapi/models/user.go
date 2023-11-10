package models

import (
	"errors"
	"fmt"
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
func UserSave(mobile string, password string) error {
	var user []User

	count, _ := orm.NewOrm().Raw("select * from user where mobile = ? limit 1", mobile).QueryRows(&user)
	if count != 0 {
		return errors.New("手机号已被使用")
	}

	result, err := orm.NewOrm().Raw("insert into user (mobile,password) values (?,?)", mobile, password).Exec()
	if err != nil {
		logs.Error(fmt.Errorf("保存用户信息异常:%w", err))
		return errors.New("内部异常")
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		logs.Error(fmt.Errorf("保存用户信息失败,用户信息未发生变化"))
		return errors.New("内部异常")
	}
	return nil
}

// 获取用户信息
func GetUserinfo(param map[string]interface{}) ([]User, int64, error) {
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

	count, err := orm.NewOrm().Raw(sql, args...).QueryRows(&user)
	if err != nil {
		logs.Error(err)
		return user, 0, errors.New("内部异常")
	}
	return user, count, nil
}
