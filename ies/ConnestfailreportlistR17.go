package ies

// ConnEstFailReportList-r17 ::= SEQUENCE OF ConnEstFailReport-r16
// SIZE (1..maxCEFReport-r17)
type ConnestfailreportlistR17 struct {
	Value []ConnestfailreportR16 `lb:1,ub:maxCEFReportR17`
}
