package ies

// MeasResultListExtIdle-r16 ::= SEQUENCE OF MeasResultIdleListEUTRA-r15
// SIZE (1..maxIdleMeasCarriersExt-r16)
type MeasresultlistextidleR16 struct {
	Value []MeasresultidlelisteutraR15 `lb:1,ub:maxIdleMeasCarriersExtR16`
}
