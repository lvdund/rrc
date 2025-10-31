package ies

// VarMeasReportList ::= SEQUENCE OF VarMeasReport
// SIZE (1..maxMeasId)
type Varmeasreportlist struct {
	Value []Varmeasreport `lb:1,ub:maxMeasId`
}
