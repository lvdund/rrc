package ies

// UAC-BarringInfoSetList-r15 ::= SEQUENCE OF UAC-BarringInfoSet-r15
// SIZE (1..maxBarringInfoSet-r15)
type UacBarringinfosetlistR15 struct {
	Value []UacBarringinfosetR15 `lb:1,ub:maxBarringInfoSetR15`
}
