package controllers

import(
//	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/context"
//	models "cannabot/cannabot-lib/models"
)

type WebOperator struct {
	beego.Controller
}

func (wo *WebOperator) Get(/*r *Report*/) {
	wo.Data["Website"] = "gmail.com"
	wo.Data["Email"] = "Matthew"
	wo.TplName = "index.tpl"
	wo.SendInstructions()
	fmt.Println("Get")
}

func (mo *WebOperator) SendInstructions() {
//		ctx := context.Context{Request: req, ResponseWriter: rw}
//		out := context.NewOutput()

//		out.JSON("{asdf: asdfasdf}", true, false)
//		report := models.InitTestReport()
//		out.JSON(report, true, false)

//		mo.Data["json"] = &json.Marshal(report)
//		mo.ServeJSON()
	fmt.Println("SendReport()")
}