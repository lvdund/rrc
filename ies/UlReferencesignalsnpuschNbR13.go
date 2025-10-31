package ies

import "rrc/utils"

// UL-ReferenceSignalsNPUSCH-NB-r13 ::= SEQUENCE
type UlReferencesignalsnpuschNbR13 struct {
	GrouphoppingenabledR13   utils.BOOLEAN
	GroupassignmentnpuschR13 utils.INTEGER `lb:0,ub:29`
}
