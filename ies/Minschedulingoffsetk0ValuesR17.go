package ies

import "rrc/utils"

// MinSchedulingOffsetK0-Values-r17 ::= SEQUENCE OF utils.INTEGER // SIZE (1..maxNrOfMinSchedulingOffsetValues-r16)
type Minschedulingoffsetk0ValuesR17 struct {
	Value []utils.INTEGER `lb:1,ub:maxNrOfMinSchedulingOffsetValuesR16`
}
