package ies

// UAC-BarringInfoSetList-v1700 ::= SEQUENCE OF UAC-BarringInfoSet-v1700
// SIZE (1..maxBarringInfoSet)
type UacBarringinfosetlistV1700 struct {
	Value []UacBarringinfosetV1700 `lb:1,ub:maxBarringInfoSet`
}
