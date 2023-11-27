package controllers

import (
	"fmt"
	"fyoukuapi/models"
	"strings"
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
	err = models.UserSave(mobile, Md5V(password))
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

	user, count, err := models.GetUserinfo(param)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	if count == 0 {
		this.JsonResult(1, "账号或密码错误")
	}
	this.JsonResult(0, "登录成功", user)
}

// 发送通知消息
func (this *UserController) UserSendMessage() {
	var (
		uids    string
		content string
	)

	uids = this.GetString("uids", "")
	if uids == "" {
		this.JsonResult(1, "接收人不能为空")
	}
	content = this.GetString("content", "")
	if content == "" {
		this.JsonResult(1, "发送内容不能为空")
	}
	arr := strings.Split(uids, ",")
	err := models.UserSendMessage(arr, content)
	if err != nil {
		this.JsonResult(1, err.Error())
	}
	this.JsonResult(0, "发送成功")
}
