package ies

import "rrc/utils"

// SRS-PosResource-r16-resourceType-r16-aperiodic-r16 ::= SEQUENCE
// Extensible
type SrsPosresourceR16ResourcetypeR16AperiodicR16 struct {
	SlotoffsetR16 *utils.INTEGER `lb:0,ub:32`
}
