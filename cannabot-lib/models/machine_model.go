// The idea here would there would be at least 1 machine
// and then an 'Operator' Model (controllers ect) can operate
// the machines
// The machines will have at least 1 sensor and at least 1 actuator

package models

import ()

type Machine_Base struct {
	Name		string          `json:"name"`
	Sensors     	[]Base		`json:"sensors"`
	Actuators   	[]Actuator   	`json:"actuators"`
	Buttons 	[]Button	`json:"buttons"`
	Indicators	[]Indicator	`json:"indicators"`
}

type Machine struct {
	Machine_Base
}

func NewMachine(name string) *Machine {
	ret := &Machine{}
	ret.Name = name
	return ret
}

func (m *Machine) AddSensor(s Sensor) {
	m.Sensors = append(m.Sensors, s)
}

func (m *Machine) AddActuator(a Actuator) {
	m.Actuators = append(m.Actuators, a)
}

func (m *Machine) AddButton(b Button) {
	m.Buttons = append(m.Buttons, b)
}

func (m *Machine) AddIndicator(i Indicator) {
	m.Indicators = append(m.Indicators, i)
}

func InitTestMachine() *Machine {
	machine := NewMachine("Test Machine")

	machine.AddSensor(NewPSISensor("Test PSI Sensor"))
	machine.AddSensor(NewTempSensor("Test Temperature Sensor"))
	machine.AddActuator(NewOnOffValveActuator("Test OnOff Valve"))
	machine.AddButton(NewToggleButton("Test Toggle Button"))
	machine.AddButton(NewPushButton("Test Push Button"))
	machine.AddIndicator(NewLEDIndicator("Test LED"))

	return machine
}
