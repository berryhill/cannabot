package controllers

import (
//	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/context"
	"cannabot/cannabot-lib/models"
	"net/http"
)

type MachineOperator struct {
	beego.Controller
	Instructions    	[]models.Instruction   	`json:"instructions"`
	CurrentReport		*models.Report 		`json:"report"`
	Machine 		*models.Machine		`json:"machine"`
	Actions			*models.Action		`json:"actions"`
}

// VVVV TOTAL SHITSHOW VVVV

func (mo *MachineOperator) Get() {
	response, err := http.Get("http://localhost:8080/test")

	if err != nil {
		fmt.Println(response)
	}

	var reportList []*models.Report

	for k := 0; k < 5; k++ {
		report := models.InitTestReport()
		report.Id = k
		reportList = append(reportList, report)
	}
	fmt.Println(reportList)

	machine := *models.InitTestMachine()
	mo.Data["json"] = &machine
	mo.ServeJSON()
//	mo.TplName = "index.tpl"
}

func (mo *MachineOperator) SendReport(/*r *Report*/) {
//	ctx := context.Context{Request: req, ResponseWriter: rw}
//	out := context.NewOutput()

//	out.JSON(report, true, false)

//	report := models.InitTestReport()

//	mo.Data["json"] = &report
//	mo.ServeJSON()
//	mo.TplName = "index.tpl"

	fmt.Println("Hello")
}
