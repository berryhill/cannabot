// The idea here would there would be at least 1 machine
// and then an 'Operator' Model (controllers ect) can operate
// the machines
// The machines will have at least 1 sensor and at least 1 actuator

package models

import ()

type Machine_Base struct {
	Name        	string          	`json:"name"`
	Sensors     	[]Sensor	   	`json:"sensors"`
	Actuators   	[]Actuator   		`json:"actuators"`
	Buttons 	[]Button		`json:"buttons"`
	Indicators	[]Indicator		`json:"indicators"`
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

func InitTestMachine() *Machine {
	machine := new(Machine)
	machine.Name = "Test Machine"

	machine.Sensors = append(machine.Sensors, NewPSISensor("Test PSI Sensor"))
	machine.Sensors = append(machine.Sensors, NewTempSensor("Test Temperature Sensor"))
	machine.Actuators = append(machine.Actuators, NewOnOffValveActuator("Test OnOff Valve"))
	machine.Buttons = append(machine.Buttons, NewToggleButton("Test Toggle Button"))
	machine.Indicators = append(machine.Indicators, NewLEDIndicator("Test LED"))

	return machine
}
