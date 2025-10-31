package ies

// ThresholdNR-r15 ::= CHOICE
const (
	ThresholdnrR15ChoiceNothing = iota
	ThresholdnrR15ChoiceNrRsrpR15
	ThresholdnrR15ChoiceNrRsrqR15
	ThresholdnrR15ChoiceNrSinrR15
)

type ThresholdnrR15 struct {
	Choice    uint64
	NrRsrpR15 *RsrpRangenrR15
	NrRsrqR15 *RsrqRangenrR15
	NrSinrR15 *RsSinrRangenrR15
}
