package controllers

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/astaxie/beego"
)

type ReadController struct {
	beego.Controller
}

//上传下载文件的页面
func (c *ReadController) Get() {
	
	c.TplName = "filerw.html"
}
//创建文件
type CreateController struct {
	beego.Controller
}
func (c *CreateController) Post(){
     //创建文件
	file, error := ioutil.TempFile("static/txtfile", "tmp")
	//文件关闭
	defer file.Close()
	if error != nil {
	   fmt.Println("创建文件失败")
	}
	c.Data["json"]=map[string]interface{}{"data":file.Name()};
	c.ServeJSON();
}
//写入文件
type WriteController struct {
	beego.Controller
}
func (c *WriteController) Post() { 
	confPath := c.GetString("path")
	info:=c.GetString("info")
	content,err := parseWriteConfig(confPath,info)
	if err != nil {
		fmt.Println(err)
    }
	fmt.Println(content)
	c.Data["json"]=map[string]interface{}{"data":string(content)};
	c.ServeJSON();
}

//写入text文件内容  
func parseWriteConfig(confPath,info string) ([]byte,error) {
	fl, err := os.OpenFile(confPath, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	defer fl.Close()
	byteinfo:=[]byte (info)
	n, err := fl.Write(byteinfo)
	if err == nil && n < len(byteinfo) {
		fmt.Println("写入失败")
		fmt.Println(err)
	}
	return byteinfo, err
}

//读取文件内容
func (c *ReadController) Post() { 
	confPath := c.GetString("path")
	fmt.Println("文件的地址:")
	fmt.Println(confPath)
	content,err := ReadFile(confPath)
	if err != nil {
		c.Data["data"]="";
		fmt.Println(err) 
    } else{
		c.Data["data"]=content;
	}
	fmt.Println(content)
	c.Data["json"]=map[string]interface{}{"data":content};
	c.ServeJSON();
}
//解析text文件内容  
func ReadFile(path string) (str string, err error) {
	//打开文件的路径
	fi, err := os.Open(path)
	if err!=nil{
		fmt.Println("打开文件失败")
		fmt.Println(err)
	}
	defer fi.Close()
	//读取文件的内容
	fd, err := ioutil.ReadAll(fi)
	if err!=nil{
		fmt.Println("读取文件失败")
		fmt.Println(err)
	}
    str = string(fd)
    return str,err
}
//删除文件
type DeleteController struct {
	beego.Controller
}
func (c *DeleteController) Post(){
	isdel:=false;
	file:=c.GetString("path");           //源文件路径
    err := os.Remove(file)               //删除文件
    if err != nil {
        //删除失败,输出错误详细信息
        fmt.Println(err)
	}else {
        //如果删除成功则输出 
		isdel=true
    }
	c.Data["json"]=map[string]interface{}{"data":isdel};
	c.ServeJSON();
}