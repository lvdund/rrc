package ies

// InterFreqCarrierFreqListExt-v1530 ::= SEQUENCE OF InterFreqCarrierFreqInfo-v1530
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistextV1530 struct {
	Value []InterfreqcarrierfreqinfoV1530 `lb:1,ub:maxFreq`
}
