package ies

import "rrc/utils"

// UAC-BarringPerCatList-r15 ::= SEQUENCE OF UAC-BarringPerCat-r15
// SIZE (1..maxAccessCat-1-r15)
type UacBarringpercatlistR15 struct {
	Value utils.Sequence[UacBarringpercatR15]
}
