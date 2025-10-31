package ies

// CRS-AssistanceInfoList-r13 ::= SEQUENCE OF CRS-AssistanceInfo-r13
// SIZE (1..maxCellReport)
type CrsAssistanceinfolistR13 struct {
	Value []CrsAssistanceinfoR13 `lb:1,ub:maxCellReport`
}
