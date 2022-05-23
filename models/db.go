package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)

// 用户信息
type Users struct {
	Id int
	Name string `orm:"unique"`  // 用户名唯一
	Pwd string 
}



func init() {
	//设置数据库连接信息
	orm.RegisterDataBase("default", "mysql", "root:zeng0428@tcp(127.0.0.1:3306)/text?charset=utf8&loc=Local")
	// 映射modle数据
	orm.RegisterModel(new(Users))
	// 生成表，第二个false要是改成true 就会强制更新表，数据全部丢失
	orm.RunSyncdb("default", false, true)
}
