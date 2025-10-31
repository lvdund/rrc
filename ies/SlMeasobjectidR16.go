package ies

import "rrc/utils"

// SL-MeasObjectId-r16 ::= utils.INTEGER (1..maxNrofSL-ObjectId-r16)
type SlMeasobjectidR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSLObjectidR16`
}
