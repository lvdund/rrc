package ies

// InterFreqCarrierFreqList ::= SEQUENCE OF InterFreqCarrierFreqInfo
// SIZE (1..maxFreq)
type Interfreqcarrierfreqlist struct {
	Value []Interfreqcarrierfreqinfo `lb:1,ub:maxFreq`
}
