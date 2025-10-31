package ies

// LogMeasInfoList2-r10 ::= SEQUENCE OF LogMeasInfo-r10
// SIZE (1..maxLogMeas-r10)
type Logmeasinfolist2R10 struct {
	Value []LogmeasinfoR10 `lb:1,ub:maxLogMeasR10`
}
