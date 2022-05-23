package controllers

import (
	"github.com/astaxie/beego"
)

type TagsController struct {
	beego.Controller
}
//标签
func (c *TagsController) TagsGet() {
	
	c.TplName = "tags.html"
}
