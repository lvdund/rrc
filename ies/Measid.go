package ies

import "rrc/utils"

// MeasId ::= utils.INTEGER (1..maxNrofMeasId)
type Measid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofMeasId`
}
