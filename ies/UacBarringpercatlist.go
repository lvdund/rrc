package ies

// UAC-BarringPerCatList ::= SEQUENCE OF UAC-BarringPerCat
// SIZE (1..maxAccessCat-1)
type UacBarringpercatlist struct {
	Value []UacBarringpercat `lb:1,ub:maxAccessCat1`
}
