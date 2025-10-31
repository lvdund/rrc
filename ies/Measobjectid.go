package ies

import "rrc/utils"

// MeasObjectId ::= utils.INTEGER (1..maxObjectId)
type Measobjectid struct {
	Value utils.INTEGER `lb:0,ub:maxObjectId`
}
