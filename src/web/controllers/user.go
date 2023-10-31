package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"web/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) UserInfo() {
	var (
		id   int
		err  error
		msg  string
		user models.User
	)

	id, err = c.GetInt("id")
	user, err = models.GetUserInfoById(id)
	msg = "查询失败"
	if err == nil {
		msg = "查询成功,name:" + user.Name
	}
	c.Ctx.WriteString(msg)
}

func (c *UserController) UserSave() {
	var (
		name   string
		avatar string
		mobile string
		err    error
		msg    string
	)

	name = c.GetString("name")
	avatar = c.GetString("avatar")
	mobile = c.GetString("mobile")
	err = models.SaveUserInfo(name, avatar, mobile)
	msg = "保存失败"
	if err == nil {
		msg = "保存成功"
	}
	c.Ctx.WriteString(msg)
}

func (c *UserController) UserUpdate() {
	var (
		id   int
		name string
		err  error
		msg  string
	)

	id, _ = c.GetInt("id")
	name = c.GetString("name")
	err = models.UserUpdateInfo(id, name)
	msg = "修改失败"
	if err == nil {
		msg = "修改成功"
	}

	c.Ctx.WriteString(msg)
}

func (c *UserController) UserDelete() {
	var (
		id  int
		msg string
		err error
	)

	id, _ = c.GetInt("id")
	err = models.UserDeleteInfo(id)
	msg = "删除失败"
	if err == nil {
		msg = "删除成功"
	}
	c.Ctx.WriteString(msg)
}

func (c *UserController) UserList() {
	var (
		err   error
		msg   string
		users []models.User
	)

	users, err = models.GetUserList()
	msg = "无数据"
	if err == nil {
		msg = ""
		for _, v := range users {
			msg += v.Name + ","
		}
	}
	c.Ctx.WriteString(msg)
}

func (c *UserController) UserInfoSql() {
	var (
		id   int
		msg  string
		err  error
		user models.User
	)

	id, err = c.GetInt("id")
	user, err = models.SqlQuery(id)
	if err == nil {
		msg = "name:" + user.Name
	} else {
		msg = "无数据"
	}
	c.Ctx.WriteString(msg)
}

func (c *UserController) UserSaveSql() {
	var (
		name   string
		mobile string
		avatar string
		msg    string
		err    error
	)

	name = c.GetString("name")
	mobile = c.GetString("mobile")
	avatar = c.GetString("avatar")

	err = models.SqlSave(name, mobile, avatar)
	msg = "保存失败"
	if err == nil {
		msg = "保存成功"
	}
	c.Ctx.WriteString(msg)
}

func (c *UserController) UserDeleteSql() {
	var (
		id  int
		msg string
		err error
	)

	id, _ = c.GetInt("id")
	err = models.SqlDelete(id)
	msg = "删除失败"
	if err == nil {
		msg = "删除成功"
	}
	c.Ctx.WriteString(msg)
}

func (c *UserController) UserUpdateSql() {
	var (
		id   int
		name string
		msg  string
		err  error
	)

	id, _ = c.GetInt("id")
	name = c.GetString("name")
	err = models.SqlUpdate(id, name)
	msg = "修改失败"
	if err == nil {
		msg = "修改成功"
	}
	c.Ctx.WriteString(msg)
}

func (c *UserController) UserListSql() {
	var (
		msg string
	)

	users, err := models.SqlList()
	msg = "无数据"
	if err == nil {
		msg = ""
		for _, v := range users {
			msg += v.Name + ","
		}
	}
	c.Ctx.WriteString(msg)
}
