package models

import ()

type Button interface {
	TypeString() string
}

type Button_Base struct {
	Name		string		`json:"name"`
	Type		string		`json:"type"`
}

type ToggleButton struct {
	Button_Base
}

func NewToggleButton(name string) *ToggleButton {
	ret := &ToggleButton{}
	ret.Name = name
	ret.Type = "toggle_button"
	return ret
}

func (led *ToggleButton) TypeString() string {
	return led.Type
}

type PushButton struct {
	Button_Base
}

func NewPushButton(name string) *PushButton {
	ret := &PushButton{}
	ret.Name = name
	ret.Type = "push_button"
	return ret
}

func (led *PushButton) TypeString() string {
	return led.Type
}

