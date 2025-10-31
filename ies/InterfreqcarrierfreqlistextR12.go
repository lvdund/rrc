package ies

// InterFreqCarrierFreqListExt-r12 ::= SEQUENCE OF InterFreqCarrierFreqInfo-r12
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistextR12 struct {
	Value []InterfreqcarrierfreqinfoR12 `lb:1,ub:maxFreq`
}
