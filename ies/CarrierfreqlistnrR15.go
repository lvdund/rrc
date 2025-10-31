package ies

// CarrierFreqListNR-r15 ::= SEQUENCE OF CarrierFreqNR-r15
// SIZE (1..maxFreq)
type CarrierfreqlistnrR15 struct {
	Value []CarrierfreqnrR15 `lb:1,ub:maxFreq`
}
