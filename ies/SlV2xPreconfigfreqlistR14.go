package ies

import "rrc/utils"

// SL-V2X-PreconfigFreqList-r14 ::= SEQUENCE OF SL-V2X-PreconfigFreqInfo-r14
// SIZE (1..maxFreqV2X-r14)
type SlV2xPreconfigfreqlistR14 struct {
	Value utils.Sequence[SlV2xPreconfigfreqinfoR14]
}
