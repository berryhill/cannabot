package routers

import (
	"cannabot/cannabot-web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test/", &controllers.WebOperator{})
	beego.Router("/task/", &controllers.WebOperator{}, "get:RequestMachine")
	//beego.Router("/task/", &controllers.WebOperator{}, "get:ListTasks;post:NewTask")
	//beego.Router("/test/:id:int", &controllers.WebOperator{}, "get:GetTask;put:UpdateTask")
}
