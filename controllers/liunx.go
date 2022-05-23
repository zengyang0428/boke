package controllers

import (

	"github.com/astaxie/beego"
)

type LIUController struct {
	beego.Controller
}
//分类
func (c *LIUController) LIuGet() {
	c.TplName = "fei/liunx.html"
}
