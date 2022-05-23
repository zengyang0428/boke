package controllers

import (
	"strings"
	"boke2/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

// 登录页面 get
func (this *RegisterController) ShowLogin() {
	this.TplName = "login.html"

}

// 登录页面 post
func (this *RegisterController) HandleLogin() {
	// 拿到浏览器数据，并去除两边的空格
	Name := strings.TrimSpace(this.GetString("userName"))
	Pwd := strings.TrimSpace(this.GetString("passWord"))
	beego.Info("账号:", Name, "密码:", Pwd)

	//数据处理
	if Name == "" || Pwd == "" {
		beego.Info("登录失败！！")
		this.TplName = "login.html"
		this.Data["errmsg"] = "登录失败！！！！！"
		return
	}
	// 查找数据库
	//获取orm对象
	o := orm.NewOrm()
	//获取查询对象
	var user models.Users
	// 查询
	user.Name = Name
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("用户名登录失败！！！")
		this.TplName = "login.html"
		this.Data["errmsg"] = "用户名登录失败！！！！！"
		return
	}
	// 判断密码是否一致
	if user.Pwd != Pwd {
		beego.Info("密码登录失败！！！")
		this.TplName = "login.html"
		this.Data["errmsg"] = "密码登录失败！！"
		return
	}
	// 测试返回视图
	//this.Ctx.WriteString("登录成功")
	// 实际情况注册成功返回到登录页面
	 this.Redirect("/login/blogs", 302)
	
}
