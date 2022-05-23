package controllers

import (
	"github.com/astaxie/beego"
)

type ArchivesController struct {
	beego.Controller
}
//归档
func (c *ArchivesController) ArchivesGet() {
	c.TplName = "archives.html"
}
