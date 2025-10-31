package ies

// NeedForNCSG-BandListNR-r17 ::= SEQUENCE OF NeedForNCSG-NR-r17
// SIZE (1..maxBands)
type NeedforncsgBandlistnrR17 struct {
	Value []NeedforncsgNrR17 `lb:1,ub:maxBands`
}
