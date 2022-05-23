package routers

import (
	"boke2/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//首页
    	beego.Router("/", &controllers.MainController{})
	//关于我
	beego.Router("/about", &controllers.AboutController{},"Get:AboutGet")
	//归档
	beego.Router("/archives", &controllers.ArchivesController{},"Get:ArchivesGet")
	//博客详情
	beego.Router("/blog", &controllers.BlogController{},"Get:BlogGet" )
	//标签
	beego.Router("/tags", &controllers.TagsController{},"Get:TagsGet")
	beego.Router("/beego", &controllers.BeeController{},"Get:BeeGet")
	beego.Router("/go", &controllers.GoController{},"Get:GOGet")
	beego.Router("/mysql", &controllers.MyController{},"Get:MGet")
	//分类
	beego.Router("/types", &controllers.TypesController{},"Get:TypesGet")
	beego.Router("/qianduan", &controllers.QianController{},"Get:QianGet")
	beego.Router("/houd", &controllers.HouController{},"Get:HouGet")
	beego.Router("/liunx", &controllers.LIUController{},"Get:LIuGet")
	beego.Router("/shujk", &controllers.ShuController{},"Get:ShuGet")
	//登陆
	beego.Router("/login", &controllers.RegisterController{}, "Get:ShowLogin;post:HandleLogin")
	//注册
	//beego.Router("/Legister", &controllers.LegisterController{},"Get:ReleaseGet;post:Registerpost")
	
	//登陆列表
	beego.Router("/login/blogs", &controllers.BlogsController{},"Get:BlogsGet")
	//登陆发布
	beego.Router("/login/release", &controllers.ReleaseController{},"Get:ReleaseGet")
	

}

