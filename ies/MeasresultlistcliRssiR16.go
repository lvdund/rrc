package ies

// MeasResultListCLI-RSSI-r16 ::= SEQUENCE OF MeasResultCLI-RSSI-r16
// SIZE (1.. maxCLI-Report-r16)
type MeasresultlistcliRssiR16 struct {
	Value []MeasresultcliRssiR16 `lb:1,ub:maxCLIReportR16`
}
