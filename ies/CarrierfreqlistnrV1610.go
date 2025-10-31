package ies

// CarrierFreqListNR-v1610 ::= SEQUENCE OF CarrierFreqNR-v1610
// SIZE (1..maxFreq)
type CarrierfreqlistnrV1610 struct {
	Value []CarrierfreqnrV1610 `lb:1,ub:maxFreq`
}
