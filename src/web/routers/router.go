package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"web/controllers"
)

func init() {
	web.Router("/", &controllers.MainController{})

	web.Router("/user/info", &controllers.UserController{}, "get:UserInfo")
	web.Router("/user/save", &controllers.UserController{}, "put:UserSave")
	web.Router("/user/update", &controllers.UserController{}, "post:UserUpdate")
	web.Router("/user/delete", &controllers.UserController{}, "delete:UserDelete")
	web.Router("/user/list", &controllers.UserController{}, "get:UserList")

	web.Router("/user/info/sql", &controllers.UserController{}, "get:UserInfoSql")
	web.Router("/user/save/sql", &controllers.UserController{}, "put:UserSaveSql")
	web.Router("/user/delete/sql", &controllers.UserController{}, "delete:UserDeleteSql")
	web.Router("/user/update/sql", &controllers.UserController{}, "post:UserUpdateSql")
	web.Router("/user/list/sql", &controllers.UserController{}, "get:UserListSql")

	web.Router("/user", &controllers.UserController{})

	var verifyFunc = func(ctx *context.Context) {
		var txt string
		txt = "禁止访问"
		ctx.WriteString(txt)
	}
	web.InsertFilter("/admin/*", web.BeforeRouter, verifyFunc)
}
