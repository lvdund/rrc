package ies

import "rrc/utils"

// SL-InterFreqInfoListV2X-r14 ::= SEQUENCE OF SL-InterFreqInfoV2X-r14
// SIZE (0..maxFreqV2X-1-r14)
type SlInterfreqinfolistv2xR14 struct {
	Value utils.Sequence[SlInterfreqinfov2xR14]
}
