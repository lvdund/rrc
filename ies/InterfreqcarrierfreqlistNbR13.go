package ies

// InterFreqCarrierFreqList-NB-r13 ::= SEQUENCE OF InterFreqCarrierFreqInfo-NB-r13
// SIZE (1..maxFreq)
type InterfreqcarrierfreqlistNbR13 struct {
	Value []InterfreqcarrierfreqinfoNbR13 `lb:1,ub:maxFreq`
}
