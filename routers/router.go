package routers

import (
	"github.com/astaxie/beego"
	"sfs/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.PointController{})
	beego.Router("/point/SavePoints", &controllers.PointController{}, "post:SavePoints")
	beego.Router("/point/GetTypeList", &controllers.PointController{}, "get:GetTypeList")
}
