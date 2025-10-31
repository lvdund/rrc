package ies

// InterFreqCarrierFreqList-v1700 ::= SEQUENCE OF InterFreqCarrierFreqInfo-v1700
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistV1700 struct {
	Value []InterfreqcarrierfreqinfoV1700 `lb:1,ub:maxFreq`
}
