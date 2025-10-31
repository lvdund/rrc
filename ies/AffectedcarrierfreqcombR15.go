package ies

// AffectedCarrierFreqComb-r15 ::= SEQUENCE OF MeasObjectId-r13
// SIZE (1..maxServCell-r13)
type AffectedcarrierfreqcombR15 struct {
	Value []MeasobjectidR13 `lb:1,ub:maxServCellR13`
}
