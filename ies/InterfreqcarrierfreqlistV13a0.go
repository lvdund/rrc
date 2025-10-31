package ies

// InterFreqCarrierFreqList-v13a0 ::= SEQUENCE OF InterFreqCarrierFreqInfo-v1360
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistV13a0 struct {
	Value []InterfreqcarrierfreqinfoV1360 `lb:1,ub:maxFreq`
}
