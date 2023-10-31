package controllers

import (
	"fyoukuapi/models"
)

type UserController struct {
	BaseController
}

// 用户注册
func (this *UserController) UserRegister() {
	var (
		mobile   string
		password string
	)

	mobile = this.GetString("mobile")
	password = this.GetString("password")

	err := CheckMobile(mobile)
	if err != nil {
		this.JsonResult(1, err.Error())
	}

	if password == "" {
		this.JsonResult(1, "密码不能为空")
	}
	_, err = models.UserSave(mobile, Md5V(password))
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	this.JsonResult(0, "注册成功")
}

// 用户登录
func (this *UserController) UserLogin() {
	var (
		mobile   string
		password string
		err      error
	)

	mobile = this.GetString("mobile")
	password = this.GetString("password")

	err = CheckMobile(mobile)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	if password == "" {
		this.JsonResult(1, "密码不能为空")
	}

	condition := make(map[string]interface{})
	condition["mobile"] = mobile
	condition["password"] = Md5V(password)
	user, res := models.GetUserinfo(condition)
	if !res {
		this.JsonResult(1, "账号或密码错误")
	}
	data := make(map[string]interface{})
	data["id"] = user.Id
	data["name"] = user.Name
	data["mobile"] = user.Mobile
	data["avatar"] = user.Avatar
	this.JsonResult(0, "登录成功", data)
}
