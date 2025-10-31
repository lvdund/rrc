package ies

// InterFreqCarrierFreqList-v1250 ::= SEQUENCE OF InterFreqCarrierFreqInfo-v1250
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistV1250 struct {
	Value []InterfreqcarrierfreqinfoV1250 `lb:1,ub:maxFreq`
}
