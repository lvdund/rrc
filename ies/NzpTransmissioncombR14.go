package ies

import "rrc/utils"

// NZP-TransmissionComb-r14 ::= utils.INTEGER (0..2)
type NzpTransmissioncombR14 struct {
	Value utils.INTEGER `lb:0,ub:2`
}
