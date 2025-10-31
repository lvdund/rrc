package ies

// NR-CarrierList-r16 ::= SEQUENCE OF MeasIdleCarrierNR-r16
// SIZE (1..maxFreqIdle-r15)
type NrCarrierlistR16 struct {
	Value []MeasidlecarriernrR16 `lb:1,ub:maxFreqIdleR15`
}
