package controllers

//
//import (
//	"fyoukuapi/config"
//	"github.com/beego/beego/v2/server/web"
//	"strconv"
//)
//
//type CommonController struct {
//    web.Controller
//}
//
//type JsonStruct struct {
//    Code  int         `json:"code"`
//    Msg   interface{} `json:"msg"`
//    Data  interface{} `json:"data"`
//    Count int64       `json:"count"`
//}
//
//func ReturnSuccess(code int, msg interface{}, data interface{}, count int64) *JsonStruct {
//    json := &JsonStruct{Code: code, Msg: msg, Data: data, Count: count}
//    return json
//}
//
//func ReturnError(code int, msg interface{}) *JsonStruct {
//    code = "Code" + strconv.Itoa(code)
//    json := &JsonStruct{Code: code, Msg: config.Code4003}
//    return json
//}
//
////func Md5V(password string) string {
////	md5 := md5.New()
////	md5.Write([]byte(password + web.AppConfig.String("md5code")))
////	return hex.EncodeToString(md5.Sum(nil))
////}
