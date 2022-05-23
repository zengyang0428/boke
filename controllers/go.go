package controllers

import (
	"github.com/astaxie/beego"
)

type GoController struct {
	beego.Controller
}
//分类
func (c *GoController) GOGet() {
	
	c.TplName = "biao/go.html"
}
