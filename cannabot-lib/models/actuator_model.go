// Same idea as the sensor

package models

import()

type Actuator_Base struct {
        Name 		string		`json:"name"`
	Type   		string 		`json:"type"`
}

type Actuator interface {
	TypeString() string
}

type OnOffValveActuator struct {
	Actuator_Base
}

func NewOnOffValveActuator(name string) *OnOffValveActuator {
	ret := &OnOffValveActuator{}
	ret.Name = name
	ret.Type = "on_off_valve_actuator"
	return ret
}

func (oova OnOffValveActuator) TypeString() string {
	return oova.Type
}
