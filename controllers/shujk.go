package controllers

import (
	"github.com/astaxie/beego"
)

type ShuController struct {
	beego.Controller
}
//分类
func (c *ShuController) ShuGet() {
	
	c.TplName = "fei/shujk.html"
}
