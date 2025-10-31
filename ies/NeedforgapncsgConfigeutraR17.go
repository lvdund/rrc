package ies

// NeedForGapNCSG-ConfigEUTRA-r17 ::= SEQUENCE
type NeedforgapncsgConfigeutraR17 struct {
	RequestedtargetbandfilterncsgEutraR17 *[]Freqbandindicatoreutra `lb:1,ub:maxBandsEUTRA`
}
