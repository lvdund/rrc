package ies

// MBMS-InterFreqCarrierTypeList-r14 ::= SEQUENCE OF MBMS-CarrierType-r14
// SIZE (1..maxFreq)
type MbmsInterfreqcarriertypelistR14 struct {
	Value []MbmsCarriertypeR14 `lb:1,ub:maxFreq`
}
