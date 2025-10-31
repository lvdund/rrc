package ies

// LogMeasInfoList-r16 ::= SEQUENCE OF LogMeasInfo-r16
// SIZE (1..maxLogMeasReport-r16)
type LogmeasinfolistR16 struct {
	Value []LogmeasinfoR16 `lb:1,ub:maxLogMeasReportR16`
}
