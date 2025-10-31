package ies

// AffectedCarrierFreqComb-r11 ::= SEQUENCE OF MeasObjectId
// SIZE (2..maxServCell-r10)
type AffectedcarrierfreqcombR11 struct {
	Value []Measobjectid `lb:2,ub:maxServCellR10`
}
