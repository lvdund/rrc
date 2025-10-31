package ies

import "rrc/utils"

// MeasId ::= utils.INTEGER (1..maxMeasId)
type Measid struct {
	Value utils.INTEGER `lb:0,ub:maxMeasId`
}
