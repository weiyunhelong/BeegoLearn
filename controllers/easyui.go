package controllers

import (
	"secondweb/models"
	"fmt"
	"github.com/astaxie/beego"
)

type EasyUIController struct {
	beego.Controller
}

func (c *EasyUIController) Get() {
	
	c.TplName = "easyui.html"
}
type EasyUIDataController struct {
	beego.Controller
}
func (c *EasyUIDataController) Post() {
	//页数
	pageno,err:=c.GetInt("page")
	if err!=nil{
       fmt.Println(err)
	}
	//每页显示的记录数
	pagesize,err:=c.GetInt("rows")
	if err!=nil{
       fmt.Println(err)
	}
	//搜索的条件
	search:=c.GetString("search")
	userList:=models.SearchDataList(pagesize,pageno,search)
	listnum:=models.GetRecordNum(search)
	c.Data["json"]=map[string]interface{}{"total":listnum,"rows":userList};
	c.ServeJSON();
}

