package ies

// SL-GapFreqInfo-r13 ::= SEQUENCE
type SlGapfreqinfoR13 struct {
	CarrierfreqR13    *ArfcnValueeutraR9
	GappatternlistR13 SlGappatternlistR13
}
