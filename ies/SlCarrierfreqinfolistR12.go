package ies

import "rrc/utils"

// SL-CarrierFreqInfoList-r12 ::= SEQUENCE OF SL-CarrierFreqInfo-r12
// SIZE (1..maxFreq)
type SlCarrierfreqinfolistR12 struct {
	Value utils.Sequence[SlCarrierfreqinfoR12]
}
