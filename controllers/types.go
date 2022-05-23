package controllers

import (
	"github.com/astaxie/beego"
)

type TypesController struct {
	beego.Controller
}
//分类
func (c *TypesController) TypesGet() {
	
	c.TplName = "types.html"
}
