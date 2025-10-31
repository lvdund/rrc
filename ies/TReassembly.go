package ies

import "rrc/utils"

// T-Reassembly ::= ENUMERATED
type TReassembly struct {
	Value utils.ENUMERATED
}

const (
	TReassemblyEnumeratedNothing = iota
	TReassemblyEnumeratedMs0
	TReassemblyEnumeratedMs5
	TReassemblyEnumeratedMs10
	TReassemblyEnumeratedMs15
	TReassemblyEnumeratedMs20
	TReassemblyEnumeratedMs25
	TReassemblyEnumeratedMs30
	TReassemblyEnumeratedMs35
	TReassemblyEnumeratedMs40
	TReassemblyEnumeratedMs45
	TReassemblyEnumeratedMs50
	TReassemblyEnumeratedMs55
	TReassemblyEnumeratedMs60
	TReassemblyEnumeratedMs65
	TReassemblyEnumeratedMs70
	TReassemblyEnumeratedMs75
	TReassemblyEnumeratedMs80
	TReassemblyEnumeratedMs85
	TReassemblyEnumeratedMs90
	TReassemblyEnumeratedMs95
	TReassemblyEnumeratedMs100
	TReassemblyEnumeratedMs110
	TReassemblyEnumeratedMs120
	TReassemblyEnumeratedMs130
	TReassemblyEnumeratedMs140
	TReassemblyEnumeratedMs150
	TReassemblyEnumeratedMs160
	TReassemblyEnumeratedMs170
	TReassemblyEnumeratedMs180
	TReassemblyEnumeratedMs190
	TReassemblyEnumeratedMs200
	TReassemblyEnumeratedSpare1
)
