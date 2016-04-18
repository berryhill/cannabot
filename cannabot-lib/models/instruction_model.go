package models

import(
)

type Instruction_Base struct {
	BaseModel
}

type Instruction struct {
	Name			string			`json:"name"`
	Id 			int			`json:"id"`
	Paremeters		[]*Parameter		`json:"parameters"`
}



