package ies

// NeedForNCSG-EUTRA-r17 ::= SEQUENCE
type NeedforncsgEutraR17 struct {
	BandeutraR17     Freqbandindicatoreutra
	GapindicationR17 NeedforncsgEutraR17GapindicationR17
}
