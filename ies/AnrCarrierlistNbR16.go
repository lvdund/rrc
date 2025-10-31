package ies

// ANR-CarrierList-NB-r16 ::= SEQUENCE OF ANR-Carrier-NB-r16
// SIZE (1..maxFreqANR-NB-r16)
type AnrCarrierlistNbR16 struct {
	Value []AnrCarrierNbR16 `lb:1,ub:maxFreqANRNbR16`
}
