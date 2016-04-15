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
	fmt.Println("MachineOperator.Get")

	var ReportList []*models.Report

	for k := 0; k < 5; k++ {
		fmt.Println(k)
		report := models.InitTestReport()
		report.Id = k
		ReportList = append(ReportList, report)
	}
	fmt.Println(ReportList)

	machine := *models.InitTestMachine()

	mo.Data["json"] = &machine
	mo.ServeJSON()
//	mo.TplName = "index.tpl"
}



func (mo *MachineOperator) SendReport() {
//	ctx := context.Context{Request: req, ResponseWriter: rw}
//	out := context.NewOutput()

//	out.JSON(report, true, false)

	report := models.InitTestReport()

	mo.Data["json"] = &report
	mo.ServeJSON()
	mo.TplName = "index.tpl"

//	fmt.Println("Hello")
}