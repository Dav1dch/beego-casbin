package routers

import (
	"beego-casbin/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/api/hellworld", &controllers.MainController{}, "POST:HelloWorld")
	beego.Router("/api/helloworld", &controllers.MainController{}, "GET:HelloWorld")
}
