package ies

import "rrc/utils"

// MBMS-SessionInfoList-r9 ::= SEQUENCE OF MBMS-SessionInfo-r9
// SIZE (0..maxSessionPerPMCH)
type MbmsSessioninfolistR9 struct {
	Value []MbmsSessioninfoR9
}
