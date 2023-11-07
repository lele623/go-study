package controllers

import (
	"fmt"
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
	param := make(map[string]interface{})
	param["mobile"] = this.GetString("mobile", "")
	param["password"] = this.GetString("password", "")

	err := CheckMobile(fmt.Sprint(param["mobile"]))
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	if param["password"] == "" {
		this.JsonResult(1, "密码不能为空")
	}
	param["password"] = Md5V(fmt.Sprint(param["password"]))

	user, num := models.GetUserinfo(param)
	if num == 0 {
		this.JsonResult(1, "账号或密码错误")
	}
	this.JsonResult(0, "登录成功", user)
}
