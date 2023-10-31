package controllers

import (
	"fyoukuapi/models"
	"regexp"
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

	if mobile == "" {
		this.JsonResult(1, "手机号不能为空")
	}
	result, _ := regexp.MatchString(`^1[3456789]\d{9}$`, mobile)
	if !result {
		this.JsonResult(1, "手机号格式不正确")
	}
	if password == "" {
		this.JsonResult(1, "手机号格式不正确")
	}
	_, err := models.UserSave(mobile, password)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	this.JsonResult(0, "注册成功")
}
