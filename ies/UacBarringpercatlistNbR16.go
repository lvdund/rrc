package ies

// UAC-BarringPerCatList-NB-r16 ::= SEQUENCE OF UAC-BarringPerCat-NB-r16
// SIZE (1..maxAccessCat-1-r15)
type UacBarringpercatlistNbR16 struct {
	Value []UacBarringpercatNbR16 `lb:1,ub:maxAccessCat1R15`
}
