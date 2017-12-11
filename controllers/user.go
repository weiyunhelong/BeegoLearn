package controllers

import (
	"BeegoLearn/models"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	
	c.TplName = "list.html"
}
func (c *UserController) Post() {
	pageno,err:=c.GetInt("pageno")
	if err!=nil{
       fmt.Println(err)
	}
	search:=c.GetString("search")
	userList:=models.SearchDataList(3,pageno,search)
	listnum:=models.GetRecordNum(search)
	c.Data["json"]=map[string]interface{}{"Count":listnum,"PageSize":3,"Page":pageno,"DataList":userList};
	c.ServeJSON();
}

type YonghuController struct {
	beego.Controller
}
func (c *YonghuController) Post() {
	pageno,err:=c.GetInt("pageno")
	if err!=nil{
       fmt.Println(err)
	}
	search:=c.GetString("search")
	userList:=models.SearchDataList(3,pageno,search)
	listnum:=models.GetRecordNum(search)
	c.Data["json"]=map[string]interface{}{"Count":listnum,"PageSize":3,"Page":pageno,"DataList":userList};
	c.ServeJSON();
}
