package routers

import (
	"BeegoLearn/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//默认的
	beego.Router("/", &controllers.MainController{})
	//JS分页
	beego.Router("/Home/PageData", &controllers.UserController{})
	beego.Router("/Home/PageNextData", &controllers.YonghuController{})
	//Bootstrap运用
	beego.Router("/Home/Index", &controllers.PageController{})
	//Easyui使用
	beego.Router("/Home/EasyUI", &controllers.EasyUIController{})
	beego.Router("/Home/EasyUIData", &controllers.EasyUIDataController{})
	//文件的上传下载
	beego.Router("/Home/FileOpt", &controllers.FileOptUploadController{})
	beego.Router("/Home/FileDown", &controllers.FileOptDownloadController{})
	//文件的创建，写入，读取，删除
	beego.Router("/Home/FileRead", &controllers.ReadController{})
	beego.Router("/Home/FileWrite", &controllers.WriteController{})
	beego.Router("/Home/FileCreate", &controllers.CreateController{})
	beego.Router("/Home/FileDelete", &controllers.DeleteController{})
	//Api接口部分
	beego.Router("/api/Html", &controllers.ApiController{})
	beego.Router("/api/GetJson", &controllers.ApiJsonController{})
	beego.Router("/api/GetXml", &controllers.ApiXMLController{})
	beego.Router("/api/GetJsonp", &controllers.ApiJsonpController{})
	beego.Router("/api/GetDictionary", &controllers.ApiDictionaryController{})
	beego.Router("/api/GetParams", &controllers.ApiParamsController{})
	//session部分
	beego.Router("/Home/Login", &controllers.LoginController{})
	beego.Router("/Home/Logout", &controllers.LogoutController{})
}
