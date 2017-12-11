package controllers

import (
	"github.com/astaxie/beego"
)

type LayoutController struct {
	beego.Controller
}
//登录页面
func (c *LayoutController) Get() {
	//布局页面
	c.Layout = "layout.html"
	//内容页面
	c.TplName = "content.html"
	//其他的部分
	c.LayoutSections = make(map[string]string)
	//页面使用的css部分
	c.LayoutSections["HtmlHead"] = "head.tpl"
	//页面使用的js部分
	c.LayoutSections["Scripts"] = "scripts.tpl"
	//页面的补充部分，可做为底部的版权部分时候
    c.LayoutSections["SideBar"] = "sidebar.tpl"
}