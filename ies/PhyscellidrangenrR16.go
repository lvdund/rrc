package ies

import "rrc/utils"

// PhysCellIdRangeNR-r16 ::= SEQUENCE
type PhyscellidrangenrR16 struct {
	Start PhyscellidnrR15
	Range *utils.ENUMERATED
}
