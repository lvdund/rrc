package ies

// InterFreqCarrierFreqList-v1310 ::= SEQUENCE OF InterFreqCarrierFreqInfo-v1310
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistV1310 struct {
	Value []InterfreqcarrierfreqinfoV1310 `lb:1,ub:maxFreq`
}
