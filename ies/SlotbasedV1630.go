package ies

import "rrc/utils"

// SlotBased-v1630 ::= SEQUENCE
type SlotbasedV1630 struct {
	TcimappingR16          SlotbasedV1630TcimappingR16
	SequenceoffsetforrvR16 utils.INTEGER `lb:0,ub:0`
}
