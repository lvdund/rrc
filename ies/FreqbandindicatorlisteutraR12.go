package ies

// FreqBandIndicatorListEUTRA-r12 ::= SEQUENCE OF FreqBandIndicator-r11
// SIZE (1..maxBands)
type FreqbandindicatorlisteutraR12 struct {
	Value []FreqbandindicatorR11 `lb:1,ub:maxBands`
}
