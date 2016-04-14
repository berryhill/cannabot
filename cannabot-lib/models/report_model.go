package models

import (
)

type Report struct {
Id		string		`json:"id`
//	RawData		string		`json:"raw_data"`
}

/*
func SendReport(r Report) {
	//TODO Implement
}
*/

func InitTestReport() *Report {
report := new(Report)
report.Id = "Test Report"
return report
}