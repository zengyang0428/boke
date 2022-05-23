package controllers

import (
	"github.com/astaxie/beego"
)

type MyController struct {
	beego.Controller
}
//分类
func (c *MyController) MGet() {
	
	c.TplName = "biao/mysql.html"
}
