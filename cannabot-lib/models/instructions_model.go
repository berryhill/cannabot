package models

import ()

type Instructions struct {
	BaseModel
	InstructionList		[]*Instruction		`json:"instructions"`
}

func (i *Instructions) GetInstructionByIndex(index int) *Instruction {
	return i.InstructionList[index]
}

func (i *Instructions) AddInstruction(instruction *Instruction) {
	i.InstructionList = append(i.InstructionList, instruction)
}

func (i *Instructions) ReplaceInstructionByIndex(index int, instruction *Instruction) {
	//TODO implement
}

