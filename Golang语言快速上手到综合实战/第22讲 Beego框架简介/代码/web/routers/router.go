package routers

import (
	"WEB/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "get:Get;post:Post")
    beego.Router("/test", &controllers.TestController{}, "get:Get;post:Post")
}
