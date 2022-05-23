package main

import (
	_ "boke2/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	beego.Run()
}
