package models

import (
)

type Report struct {
	Name		string		`json:"name"`
	Id		int		`json:"id"`
//	RawData		string		`json:"raw_data"`
	Machine 	*Machine	`json:"machine"`
}

/*
func SendReport(r Report) {
	//TODO Implement
}
*/

func InitTestReport() *Report {
	report := new(Report)
	report.Name = "Test Report"
	report.Id = -1
	return report
}

