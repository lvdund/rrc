package ies

import "rrc/utils"

// MeasGapId-r17 ::= utils.INTEGER (1..maxNrofGapId-r17)
type MeasgapidR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofGapIdR17`
}
