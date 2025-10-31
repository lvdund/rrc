package ies

// InterFreqCarrierFreqListExt-v1360 ::= SEQUENCE OF InterFreqCarrierFreqInfo-v1360
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistextV1360 struct {
	Value []InterfreqcarrierfreqinfoV1360 `lb:1,ub:maxFreq`
}
