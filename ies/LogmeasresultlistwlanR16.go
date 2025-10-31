package ies

// LogMeasResultListWLAN-r16 ::= SEQUENCE OF LogMeasResultWLAN-r16
// SIZE (1..maxWLAN-Id-Report-r16)
type LogmeasresultlistwlanR16 struct {
	Value []LogmeasresultwlanR16 `lb:1,ub:maxWLANIdReportR16`
}
