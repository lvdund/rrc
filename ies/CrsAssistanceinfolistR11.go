package ies

// CRS-AssistanceInfoList-r11 ::= SEQUENCE OF CRS-AssistanceInfo-r11
// SIZE (1..maxCellReport)
type CrsAssistanceinfolistR11 struct {
	Value []CrsAssistanceinfoR11 `lb:1,ub:maxCellReport`
}
