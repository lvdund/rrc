package ies

// VarMeasReportList ::= SEQUENCE OF VarMeasReport
// SIZE (1..maxNrofMeasId)
type Varmeasreportlist struct {
	Value []Varmeasreport `lb:1,ub:maxNrofMeasId`
}
