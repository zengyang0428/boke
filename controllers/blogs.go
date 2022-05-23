package controllers

import (
	"github.com/astaxie/beego"
)

type BlogsController struct {
	beego.Controller
}

//登陆发布
func (c *BlogsController) BlogsGet() {
	c.TplName = "blogs.html"
}
func (c *BlogsController) Blogsxiugai() {

}