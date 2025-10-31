package ies

import "rrc/utils"

// CFRA-SSB-Resource ::= SEQUENCE
// Extensible
type CfraSsbResource struct {
	Ssb                       SsbIndex
	RaPreambleindex           utils.INTEGER  `lb:0,ub:63`
	MsgaPuschResourceIndexR16 *utils.INTEGER `lb:0,ub:3071,ext`
}
