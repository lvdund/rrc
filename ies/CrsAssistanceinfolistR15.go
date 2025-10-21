package ies

import "rrc/utils"

// CRS-AssistanceInfoList-r15 ::= SEQUENCE OF CRS-AssistanceInfo-r15
// SIZE (1..maxCellReport)
type CrsAssistanceinfolistR15 struct {
	Value []CrsAssistanceinfoR15
}
