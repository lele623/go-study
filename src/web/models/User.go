package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Id      int
	Name    string
	AddTime int64
	Status  int
	Mobile  string
	Avatar  string
}

func init() {
	orm.RegisterModel(new(User))
}

// 根据id查询
func GetUserInfoById(id int) (User, error) {
	var (
		err  error
		user User
	)
	o := orm.NewOrm()
	user = User{Id: id}
	err = o.Read(&user)
	return user, err
}

// 添加用户信息
func SaveUserInfo(name string, avatar string, mobile string) error {
	var (
		err  error
		user User
	)
	user.Name = name
	user.Avatar = avatar
	user.Mobile = mobile
	user.Status = 0
	o := orm.NewOrm()
	_, err = o.Insert(&user)
	return err
}

// 修改用户信息
func UserUpdateInfo(id int, name string) error {
	var (
		err error
	)

	user := User{Id: id}
	o := orm.NewOrm()
	err = o.Read(&user)
	if o.Read(&user) == nil {
		_, err = o.Update(&user)
	}
	return err
}

// 删除用户信息
func UserDeleteInfo(id int) error {
	var (
		err error
	)

	user := User{Id: id}
	o := orm.NewOrm()
	err = o.Read(&user)
	if err == nil {
		_, err = o.Delete(&user)
	}
	return err
}

// 获取用户列表
func GetUserList() ([]User, error) {
	var (
		users []User
		err   error
	)

	o := orm.NewOrm()
	_, err = o.QueryTable("user").Filter("id__gt", 10).Limit(3).All(&users, "id", "name")
	if err != nil {
		return nil, err
	}
	return users, err
}

// sql原生查询
func SqlQuery(id int) (User, error) {
	var (
		user User
		err  error
	)

	o := orm.NewOrm()
	err = o.Raw("select `name`,`mobile` from user where id=? limit 1", id).QueryRow(&user)
	return user, err
}

// sql原生添加
func SqlSave(name string, mobile string, avatar string) error {
	var (
		err error
	)

	o := orm.NewOrm()
	_, err = o.Raw("insert into user (name,mobile,avatar,status) values(?,?,?,?)", name, mobile, avatar, 0).Exec()
	return err
}

// sql原生删除
func SqlDelete(id int) error {
	var (
		err error
	)

	o := orm.NewOrm()
	_, err = o.Raw("delete from user where id=?", id).Exec()
	return err
}

// sql修改
func SqlUpdate(id int, name string) error {
	var (
		err error
	)

	o := orm.NewOrm()
	_, err = o.Raw("update user set name=? where id=?", name, id).Exec()
	return err
}

// sql列表
func SqlList() ([]User, error) {
	var (
		users []User
		err   error
	)

	o := orm.NewOrm()
	_, err = o.Raw("select id,name from user where id >10 limit 4").QueryRows(&users)
	return users, err
}
