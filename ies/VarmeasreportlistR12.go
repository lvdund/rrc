package ies

// VarMeasReportList-r12 ::= SEQUENCE OF VarMeasReport
// SIZE (1..maxMeasId-r12)
type VarmeasreportlistR12 struct {
	Value []Varmeasreport `lb:1,ub:maxMeasIdR12`
}
