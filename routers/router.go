package routers

import (
	"secondweb/controllers"
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
}
