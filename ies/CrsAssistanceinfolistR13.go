package ies

import "rrc/utils"

// CRS-AssistanceInfoList-r13 ::= SEQUENCE OF CRS-AssistanceInfo-r13
// SIZE (1..maxCellReport)
type CrsAssistanceinfolistR13 struct {
	Value []CrsAssistanceinfoR13
}
