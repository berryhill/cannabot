// Same idea as the sensor

package models

import()

type Actuator_Base struct {
    Type        string      `json:"type"`
}

type Actuator interface {
    TypeString()
}

type OnOffValveActuator struct {
    Actuator_Base
}

func (oova OnOffValveActuator) NewOnOffValveActuator() *OnOffValveActuator {
    ret := &OnOffValveActuator{}
    ret.Type = "on_off_valve_actuator"
    return ret
}

func (oova OnOffValveActuator) TypeString() string {
    return oova.Type
}


