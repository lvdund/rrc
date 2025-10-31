package ies

import "rrc/utils"

// MeasObjectId-v1310 ::= utils.INTEGER (maxObjectId-Plus1-r13..maxObjectId-r13)
type MeasobjectidV1310 struct {
	Value utils.INTEGER `lb:maxObjectIdPlus1R13,ub:maxObjectIdR13`
}
