package ies

// CarrierFreqsListGERAN-NB-r16 ::= SEQUENCE OF CarrierFreqsGERAN-NB-r16
// SIZE (1..maxFreqsGERAN-NB-r16)
type CarrierfreqslistgeranNbR16 struct {
	Value []CarrierfreqsgeranNbR16 `lb:1,ub:maxFreqsGERANNbR16`
}
