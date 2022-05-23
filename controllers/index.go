package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
//首页
func (c *MainController) Get() {
	c.TplName = "index.html"
}
