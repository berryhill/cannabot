package models

import (
)

type Report struct {
	Name		string		`json:"name`
	Id		int		`json:"id`
//	RawData		string		`json:"raw_data"`
}

/*
func SendReport(r Report) {
	//TODO Implement
}
*/


func (r *Report)addName(n string) {
	//TODO Implement
}

func InitTestReport() *Report {
	report := new(Report)
	report.Name = "Test Report"
	report.Id = -1
	return report
}
