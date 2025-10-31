package ies

import "rrc/utils"

// SlotBased-r16 ::= SEQUENCE
type SlotbasedR16 struct {
	TcimappingR16          SlotbasedR16TcimappingR16
	SequenceoffsetforrvR16 utils.INTEGER `lb:0,ub:3`
}
