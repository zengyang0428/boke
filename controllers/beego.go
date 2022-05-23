package controllers

import (
	"github.com/astaxie/beego"
)

type BeeController struct {
	beego.Controller
}
//分类
func (c *BeeController) BeeGet() {
	
	c.TplName = "biao/beego.html"
}
