package ies

// MeasResultListWLAN-r14 ::= SEQUENCE OF MeasResultWLAN-r13
// SIZE (1..maxWLAN-Id-Report-r14)
type MeasresultlistwlanR14 struct {
	Value []MeasresultwlanR13 `lb:1,ub:maxWLANIdReportR14`
}
