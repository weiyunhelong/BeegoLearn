package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}
//登录页面
func (c *LoginController) Get() {
	c.TplName = "login.html"
}
//登录功能
func (c *LoginController) Post() {
	name:=c.GetString("name")
	pwd:=c.GetString("pwd")
	islogin:=0
	if name=="admin" && pwd=="123456"{
		c.SetSession("loginuser", "adminuser")
		fmt.Println("当前的session:")
		fmt.Println(c.CruSession)
	}else if name!="admin"{
		islogin=1
	}else if pwd!="123456"{
		islogin=2
	}
	c.Data["json"]=map[string]interface{}{"islogin":islogin};
	c.ServeJSON();
}
//退出
type LogoutController struct {
	beego.Controller
}
//登录退出功能
func (c *LogoutController) Post() {
	v := c.GetSession("loginuser")
	islogin:=false
	if v != nil {
	  //删除指定的session	
	  c.DelSession("loginuser")
	  //销毁全部的session
	  c.DestroySession()
	  islogin=true
	  
	 fmt.Println("当前的session:")
	 fmt.Println(c.CruSession)
	}
	c.Data["json"]=map[string]interface{}{"islogin":islogin};
	c.ServeJSON();
}