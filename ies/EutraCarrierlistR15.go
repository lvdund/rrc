package ies

// EUTRA-CarrierList-r15 ::= SEQUENCE OF MeasIdleCarrierEUTRA-r15
// SIZE (1..maxFreqIdle-r15)
type EutraCarrierlistR15 struct {
	Value []MeasidlecarriereutraR15 `lb:1,ub:maxFreqIdleR15`
}
