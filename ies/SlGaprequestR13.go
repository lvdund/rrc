package ies

import "rrc/utils"

// SL-GapRequest-r13 ::= SEQUENCE OF SL-GapFreqInfo-r13
// SIZE (1..maxFreq)
type SlGaprequestR13 struct {
	Value []SlGapfreqinfoR13
}
