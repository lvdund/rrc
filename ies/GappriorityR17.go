package ies

import "rrc/utils"

// GapPriority-r17 ::= utils.INTEGER (1..maxNrOfGapPri-r17)
type GappriorityR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrOfGapPriR17`
}
