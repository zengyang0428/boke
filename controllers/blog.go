package controllers

import (
	"github.com/astaxie/beego"
)

type BlogController struct {
	beego.Controller
}
//博客详情
func (c *BlogController) BlogGet() {
	c.TplName = "blog.html"
}

