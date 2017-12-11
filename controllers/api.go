package controllers

import (
	"github.com/astaxie/beego"
)
//Api页面
type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	c.TplName="api.html"
}
//JSON格式的数据
type ApiJsonController struct {
	beego.Controller
}

func (c *ApiJsonController) Get() {
	//注意此处的json，必须是json
	c.Data["json"] = "ABCDEFG"
	c.ServeJSON()
}

//XML格式的数据
type ApiXMLController struct {
	beego.Controller
}

func (c *ApiXMLController) Get() {
	//注意此处的xml，必须是xml
	c.Data["xml"] = "BCDEFGH"
	c.ServeXML()
}

//Jsonp格式的数据
type ApiJsonpController struct {
	beego.Controller
}

func (c *ApiJsonpController) Get() {
	//注意此处的jsonp，必须是jsonp
	c.Data["jsonp"] = "CDEFGHI"
	c.ServeJSONP()
}

//字典表格式的数据
type ApiDictionaryController struct {
	beego.Controller
}

func (c *ApiDictionaryController) Get() {
	c.Data["json"]=map[string]interface{}{"name":"ABC123","rows":45,"flag":true};
	c.ServeJSON()
}

//带参数的表格式的数据
type ApiParamsController struct {
	beego.Controller
}

func (c *ApiParamsController) Get() {
	search:=c.GetString("name")
	c.Data["json"]=map[string]interface{}{"name":search,"rows":45,"flag":false};
	c.ServeJSON()
}