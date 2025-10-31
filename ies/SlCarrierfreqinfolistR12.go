package ies

// SL-CarrierFreqInfoList-r12 ::= SEQUENCE OF SL-CarrierFreqInfo-r12
// SIZE (1..maxFreq)
type SlCarrierfreqinfolistR12 struct {
	Value []SlCarrierfreqinfoR12 `lb:1,ub:maxFreq`
}
