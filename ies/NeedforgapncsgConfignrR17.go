package ies

// NeedForGapNCSG-ConfigNR-r17 ::= SEQUENCE
type NeedforgapncsgConfignrR17 struct {
	RequestedtargetbandfilterncsgNrR17 *[]Freqbandindicatornr `lb:1,ub:maxBands`
}
