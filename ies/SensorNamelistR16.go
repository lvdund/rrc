package ies

import "rrc/utils"

// Sensor-NameList-r16 ::= SEQUENCE
type SensorNamelistR16 struct {
	MeasuncombarpreR16 *utils.BOOLEAN
	Measuespeed        *utils.BOOLEAN
	Measueorientation  *utils.BOOLEAN
}
