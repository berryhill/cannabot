package models

import ()

type Indicator interface {
	TypeString() string
}

type Indicator_Base struct {
	Name		string		`json:"name"`
	Type        	string		`json:"type"`
}

type LEDIndicator struct {
	Indicator_Base
}

func NewLEDIndicator(name string) *LEDIndicator {
	ret := &LEDIndicator{}
	ret.Name = name
	ret.Type = "led"
	return ret
}

func (led LEDIndicator) LEDIndicator() *LEDIndicator {
	ret := &LEDIndicator{}
	ret.Type = "led"
	return ret
}

func (led LEDIndicator) TypeString() string {
	return led.Type
}


