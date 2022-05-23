package controllers

import (
	"boke2/models"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LegisterController struct {
	beego.Controller
}

//注册
func (this *LegisterController) ReleaseGet() {
	this.TplName = "Legister.html"
}

// 注册页面 post
func (this *LegisterController) Registerpost() {
	// 获取浏览器传递的值，并去除两边的空格

	Name := strings.TrimSpace(this.GetString("userName"))
	Pwd := strings.TrimSpace(this.GetString("passWord")) 
	// beego.Info("账号:", Name, "密码:", Pwd)
	// 数据处理
	if Name == "" || Pwd == "" {
		beego.Info("用户名或密码不能为空")
		this.TplName = "Legister.html"
		this.Data["errmsg"] = "用户名或密码不能为空 ！"
		return
	} 
	// 插入数据库（数据库表，Users）
	  //获取orm对象
	o := orm.NewOrm()
	//   获取插入对象
	user := models.Users{}
	//   插入数值
	user.Name = Name
	user.Pwd = Pwd

	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据失败，用户相同或者其他错误！！！")
		this.TplName = "Legister.html"
		this.Data["errmsg"] = "插入数据失败，用户相同或者其他错误！！！！"
		return
	}
	// 测试返回视图
	 this.Ctx.WriteString("注册成功！！！")
	// 实际情况注册成功返回到登录页面
	this.Redirect("/login", 302)

}
