package ies

// ThresholdListNR-r15 ::= SEQUENCE
type ThresholdlistnrR15 struct {
	NrRsrpR15 *RsrpRangenrR15
	NrRsrqR15 *RsrqRangenrR15
	NrSinrR15 *RsSinrRangenrR15
}
