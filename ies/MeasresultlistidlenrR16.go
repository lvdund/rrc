package ies

// MeasResultListIdleNR-r16 ::= SEQUENCE OF MeasResultIdleNR-r16
// SIZE (1..maxIdleMeasCarriers-r16)
type MeasresultlistidlenrR16 struct {
	Value []MeasresultidlenrR16 `lb:1,ub:maxIdleMeasCarriersR16`
}
