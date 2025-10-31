package ies

// InterFreqCarrierFreqList-v1610 ::= SEQUENCE OF InterFreqCarrierFreqInfo-v1610
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistV1610 struct {
	Value []InterfreqcarrierfreqinfoV1610 `lb:1,ub:maxFreq`
}
