package ies

import "rrc/utils"

// MinSchedulingOffsetK2-Values-r16 ::= SEQUENCE OF utils.INTEGER // SIZE (1..maxNrOfMinSchedulingOffsetValues-r16)
type Minschedulingoffsetk2ValuesR16 struct {
	Value []utils.INTEGER `lb:1,ub:maxNrOfMinSchedulingOffsetValuesR16`
}
