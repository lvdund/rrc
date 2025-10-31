package ies

// MeasResultIdleListEUTRA-r15 ::= SEQUENCE OF MeasResultIdleEUTRA-r15
// SIZE (1..maxCellMeasIdle-r15)
type MeasresultidlelisteutraR15 struct {
	Value []MeasresultidleeutraR15 `lb:1,ub:maxCellMeasIdleR15`
}
