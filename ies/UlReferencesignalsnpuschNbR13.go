package ies

import "rrc/utils"

// UL-ReferenceSignalsNPUSCH-NB-r13 ::= SEQUENCE
type UlReferencesignalsnpuschNbR13 struct {
	GrouphoppingenabledR13   bool
	GroupassignmentnpuschR13 utils.INTEGER
}
