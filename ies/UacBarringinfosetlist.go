package ies

// UAC-BarringInfoSetList ::= SEQUENCE OF UAC-BarringInfoSet
// SIZE (1..maxBarringInfoSet)
type UacBarringinfosetlist struct {
	Value []UacBarringinfoset `lb:1,ub:maxBarringInfoSet`
}
