package ies

// NeedForNCSG-NR-r17 ::= SEQUENCE
type NeedforncsgNrR17 struct {
	BandnrR17        Freqbandindicatornr
	GapindicationR17 NeedforncsgNrR17GapindicationR17
}
