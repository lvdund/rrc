package ies

// InterFreqCarrierFreqListExt-v1280 ::= SEQUENCE OF InterFreqCarrierFreqInfo-v10j0
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistextV1280 struct {
	Value []InterfreqcarrierfreqinfoV10j0 `lb:1,ub:maxFreq`
}
