package ies

// NeedForGapsNR-r16 ::= SEQUENCE
type NeedforgapsnrR16 struct {
	BandnrR16        Freqbandindicatornr
	GapindicationR16 NeedforgapsnrR16GapindicationR16
}
