package controllers

import (
	"github.com/astaxie/beego"
)

type QianController struct {
	beego.Controller
}
//分类
func (c *QianController) QianGet() {
	
	c.TplName = "fei/qianduan.html"
}
