package models

import()

type Parameter struct {
	BaseModel
	ParameterCount		int		`json:"parametercount"`
}

func (p *Parameter) GetParameterCount() int {
	return p.ParameterCount
}