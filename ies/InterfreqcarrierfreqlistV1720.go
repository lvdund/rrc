package ies

// InterFreqCarrierFreqList-v1720 ::= SEQUENCE OF InterFreqCarrierFreqInfo-v1720
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistV1720 struct {
	Value []InterfreqcarrierfreqinfoV1720 `lb:1,ub:maxFreq`
}
