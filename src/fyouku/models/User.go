package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type UserInfo struct {
	Id      int
	Name    string
	AddTime int64
	Avatar  string
}

// 判断用户名密码是否正确
func IsMobileLogin(mobile string, password string) string {
	req := httplib.Get(beego.AppConfig.String("apiurl") + "/user/login")
	//req := httplib.Post(beego.AppConfig.String("microApi") + "/fyoukuApi/user/user/LoginDo")
	req.Param("mobile", mobile)
	req.Param("password", password)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}

// 保存用户
func UserSave(mobile string, password string) string {
	req := httplib.Put(beego.AppConfig.String("apiurl") + "/user/register")
	req.Param("mobile", mobile)
	req.Param("password", password)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}

// 发送消息
func SendMessageDo(uids string, content string) string {
	req := httplib.Post(beego.AppConfig.String("apiurl") + "/user/send/message")
	req.Param("uids", uids)
	req.Param("content", content)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}
