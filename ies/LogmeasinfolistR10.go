package ies

import "rrc/utils"

// LogMeasInfoList-r10 ::= SEQUENCE OF LogMeasInfo-r10
// SIZE (1..maxLogMeasReport-r10)
type LogmeasinfolistR10 struct {
	Value []LogmeasinfoR10
}
