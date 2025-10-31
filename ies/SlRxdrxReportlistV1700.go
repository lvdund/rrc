package ies

// SL-RxDRX-ReportList-v1700 ::= SEQUENCE OF SL-RxDRX-Report-v1700
// SIZE (1..maxNrofSL-Dest-r16)
type SlRxdrxReportlistV1700 struct {
	Value []SlRxdrxReportV1700 `lb:1,ub:maxNrofSLDestR16`
}
