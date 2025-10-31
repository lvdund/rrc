package ies

// InterFreqCarrierFreqList-v1730 ::= SEQUENCE OF InterFreqCarrierFreqInfo-v1730
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistV1730 struct {
	Value []InterfreqcarrierfreqinfoV1730 `lb:1,ub:maxFreq`
}
