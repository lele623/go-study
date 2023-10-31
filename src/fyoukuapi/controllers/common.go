package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/beego/beego/v2/server/web"
	"regexp"
)

type CommonController struct {
	web.Controller
}

// md5加密
func Md5V(password string) string {
	md5New := md5.New()
	md5code, _ := web.AppConfig.String("md5code")
	md5New.Write([]byte(password + md5code))
	return hex.EncodeToString(md5New.Sum(nil))
}

// 检查手机号
func CheckMobile(mobile string) error {
	if mobile == "" {
		return errors.New("手机号不能为空")
	}
	result, _ := regexp.MatchString(`^1[3456789]\d{9}$`, mobile)
	if !result {
		return errors.New("手机号格式不正确")
	}
	return nil
}
