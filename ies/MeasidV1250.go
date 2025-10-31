package ies

import "rrc/utils"

// MeasId-v1250 ::= utils.INTEGER (maxMeasId-Plus1..maxMeasId-r12)
type MeasidV1250 struct {
	Value utils.INTEGER `lb:maxMeasIdPlus1,ub:maxMeasIdR12`
}
