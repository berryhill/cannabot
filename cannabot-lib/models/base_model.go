package models

import()

type BaseModel struct {
	Name 		string		`json:"name"`
	Type		string		`json:"type"`
	Id 		int		`json:"id"`
}

type Base interface {
	TypeString() string
}
