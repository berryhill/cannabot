package routers

import (
	"cannabot/cannabot-web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.WebOperator{})
}
