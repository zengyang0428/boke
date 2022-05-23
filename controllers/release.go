package controllers

import (
	"github.com/astaxie/beego"
)

type ReleaseController struct {
	beego.Controller
}

//登陆发布
func (c *ReleaseController) ReleaseGet() {
	c.TplName = "release.html"
}
