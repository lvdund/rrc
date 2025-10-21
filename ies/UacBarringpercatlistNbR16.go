package ies

import "rrc/utils"

// UAC-BarringPerCatList-NB-r16 ::= SEQUENCE OF UAC-BarringPerCat-NB-r16
// SIZE (1..maxAccessCat-1-r15)
type UacBarringpercatlistNbR16 struct {
	Value []UacBarringpercatNbR16
}
