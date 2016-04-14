package controllers

import (
//	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/context"
	"cannabot/cannabot-lib/models"
)

type MachineOperator struct {
	beego.Controller
	Instructions    []interface{}   `json:"instructions"`
	CurrentReport 	*models.Report 	`json:"report"`
}

func (mo *MachineOperator) Get(/*r *Report*/) {
	fmt.Println("Get")
}



func (mo *MachineOperator) SendReport() {
//	ctx := context.Context{Request: req, ResponseWriter: rw}
//	out := context.NewOutput()

//	report := models.InitTestReport()
//	out.JSON(report, true, false)

//	mo.Data["json"] = &json.Marshal(report)
//	mo.ServeJSON()

//	fmt.Println("Hello")
}