package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["IsEmail"] = 0
	pages := []struct {
		Num int
	}{{10}, {20}, {30}, {40}}
	c.Data["Pages"] = pages
	c.TplName = "index.tpl"
}

func (c *MainController) GetDemo() {
	var title string
	title = "Demo 测试"
	c.Ctx.WriteString(title)
}
