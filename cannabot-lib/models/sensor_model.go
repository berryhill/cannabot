package models

import ()

type Sensor interface {
    TypeString() string
}

type Sensor_Base struct {
    Name        string      `json:"name"`
    Type        string      `json:"type"`
}

type PSISensor struct {
    Sensor_Base
}

// PSISensor Constructor
func NewPSISensor(name string) *PSISensor {
    ret := &PSISensor{}
    ret.Name = name
    ret.Type = "psi_sensor"
    return ret
}

// Method of the PSISensor, like a method of an object.
// 'psis PSISensor' to points the functions to the PSISensor Struct
func (psis PSISensor) TypeString() string {
    return psis.Type
}

type TempSensor struct {
    Sensor_Base
}

func NewTempSensor(name string) *TempSensor {
    ret := &TempSensor{}
    ret.Name = name
    ret.Type = "temperature_sensor"
    return ret
}

func (ts TempSensor) TypeString() string {
    return ts.Type
}

// VVVVV More Types of Sensors VVVVV //




