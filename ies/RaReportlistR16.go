package ies

// RA-ReportList-r16 ::= SEQUENCE OF RA-Report-r16
// SIZE (1..maxRAReport-r16)
type RaReportlistR16 struct {
	Value []RaReportR16 `lb:1,ub:maxRAReportR16`
}
