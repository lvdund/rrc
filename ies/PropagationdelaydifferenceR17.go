package ies

import "rrc/utils"

// PropagationDelayDifference-r17 ::= SEQUENCE OF utils.INTEGER // SIZE (1..4)
type PropagationdelaydifferenceR17 struct {
	Value []utils.INTEGER `lb:1,ub:4`
}
