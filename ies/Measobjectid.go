package ies

import "rrc/utils"

// MeasObjectId ::= utils.INTEGER (1..maxNrofObjectId)
type Measobjectid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofObjectId`
}
