package ies

// InterFreqCarrierFreqListExt-v1610 ::= SEQUENCE OF InterFreqCarrierFreqInfo-v1610
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistextV1610 struct {
	Value []InterfreqcarrierfreqinfoV1610 `lb:1,ub:maxFreq`
}
