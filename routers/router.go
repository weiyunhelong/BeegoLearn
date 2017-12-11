package routers

import (
	"BeegoLearn/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/Home/PageData", &controllers.UserController{})
	beego.Router("/Home/PageNextData", &controllers.YonghuController{})
	beego.Router("/Home/Index", &controllers.PageController{})
	beego.Router("/Home/EasyUI", &controllers.EasyUIController{})
	beego.Router("/Home/EasyUIData", &controllers.EasyUIDataController{})
	beego.Router("/Home/FileOpt", &controllers.FileOptUploadController{})
	beego.Router("/Home/FileDown", &controllers.FileOptDownloadController{})
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
}
