package controllers

import (
	"github.com/astaxie/beego"
)

type HouController struct {
	beego.Controller
}
//分类
func (c *HouController) HouGet() {
	
	c.TplName = "fei/houd.html"
}
