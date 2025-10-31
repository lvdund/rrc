package ies

// MeasResultListSRS-RSRP-r16 ::= SEQUENCE OF MeasResultSRS-RSRP-r16
// SIZE (1.. maxCLI-Report-r16)
type MeasresultlistsrsRsrpR16 struct {
	Value []MeasresultsrsRsrpR16 `lb:1,ub:maxCLIReportR16`
}
