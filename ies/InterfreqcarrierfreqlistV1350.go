package ies

// InterFreqCarrierFreqList-v1350 ::= SEQUENCE OF InterFreqCarrierFreqInfo-v1350
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistV1350 struct {
	Value []InterfreqcarrierfreqinfoV1350 `lb:1,ub:maxFreq`
}
