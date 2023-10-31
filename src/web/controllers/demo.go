package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type DemoController struct {
	beego.Controller
}

// GetHello 测试
// @router /demo/hello [get]
func (c *DemoController) GetHello() {
	var title string
	title = "Demo 控制器测试"
	c.Ctx.WriteString(title)
}
