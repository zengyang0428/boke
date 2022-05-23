package controllers

import (
	"github.com/astaxie/beego"
)

type AboutController struct {
	beego.Controller
}
//关于我
func (c *AboutController) AboutGet() {
	c.TplName = "about.html"
}
