package ies

// SL-CarrierFreqInfoList-v1310 ::= SEQUENCE OF SL-CarrierFreqInfo-v1310
// SIZE (1..maxFreq)
type SlCarrierfreqinfolistV1310 struct {
	Value []SlCarrierfreqinfoV1310 `lb:1,ub:maxFreq`
}
