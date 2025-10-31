package ies

import "rrc/utils"

// MinSchedulingOffsetK0-Values-r16 ::= SEQUENCE OF utils.INTEGER // SIZE (1..maxNrOfMinSchedulingOffsetValues-r16)
type Minschedulingoffsetk0ValuesR16 struct {
	Value []utils.INTEGER `lb:1,ub:maxNrOfMinSchedulingOffsetValuesR16`
}
