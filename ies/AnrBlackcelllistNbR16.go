package ies

import "rrc/utils"

// ANR-BlackCellList-NB-r16 ::= SEQUENCE OF PhysCellId
// SIZE (1..maxCellBlack)
type AnrBlackcelllistNbR16 struct {
	Value []Physcellid
}
