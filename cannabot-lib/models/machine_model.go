// The idea here would there would be at least 1 machine
// and then an 'Operator' Model (controllers ect) can operate
// the machines
// The machines will have at least 1 sensor and at least 1 actuator

package models

import ()

type Machine_Base struct {
	Name        string          `json:"name"`
	Sensors     []interface{}   `json:"sensors"`
    	Actuators   []interface{}   `json:"actuators"`
}

type Machine struct {
    Machine_Base
}

func (m *Machine) AddSensor(s *Sensor) {
	//TODO Implement
}

func (m *Machine) AddActuator(a *Actuator) {
	//TODO Implement
}