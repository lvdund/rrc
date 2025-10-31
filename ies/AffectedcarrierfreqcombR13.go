package ies

// AffectedCarrierFreqComb-r13 ::= SEQUENCE OF MeasObjectId-r13
// SIZE (2..maxServCell-r13)
type AffectedcarrierfreqcombR13 struct {
	Value []MeasobjectidR13 `lb:2,ub:maxServCellR13`
}
