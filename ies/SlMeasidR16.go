package ies

import "rrc/utils"

// SL-MeasId-r16 ::= utils.INTEGER (1..maxNrofSL-MeasId-r16)
type SlMeasidR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSLMeasidR16`
}
