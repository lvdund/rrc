package ies

import "rrc/utils"

// MeasObjectId-r13 ::= utils.INTEGER (1..maxObjectId-r13)
type MeasobjectidR13 struct {
	Value utils.INTEGER `lb:0,ub:maxObjectIdR13`
}
