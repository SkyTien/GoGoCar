package routers

import (
	"GoGoCar/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/api/getclients", &controllers.DBController{}, "post:GetClients")
	beego.Router("/api/getdetails", &controllers.DBController{}, "post:GetDetails")
}
