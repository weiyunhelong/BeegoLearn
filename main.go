package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "BeegoLearn/routers"
    "github.com/astaxie/beego"
)


//初始化
func init(){
	dbhost := beego.AppConfig.String("dbhost")
    dbport := beego.AppConfig.String("dbport")
    dbuser := beego.AppConfig.String("dbuser")
    dbpassword := beego.AppConfig.String("dbpassword")
    db := beego.AppConfig.String("db")

    //注册mysql Driver
    orm.RegisterDriver("mysql", orm.DRMySQL)
	//构造conn连接
	//用户名:密码@tcp(url地址)/数据库
    conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8"
    //注册数据库连接
    orm.RegisterDataBase("default", "mysql", conn)

    fmt.Printf("数据库连接成功！%s\n", conn)  
}

func main() {
    o := orm.NewOrm()
    o.Using("default") // 默认使用 default，你可以指定为其他数据库
    //启用Session
    beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
