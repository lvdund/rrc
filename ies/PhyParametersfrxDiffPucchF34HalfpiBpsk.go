package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-pucch-F3-4-HalfPi-BPSK ::= ENUMERATED
type PhyParametersfrxDiffPucchF34HalfpiBpsk struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffPucchF34HalfpiBpskEnumeratedNothing = iota
	PhyParametersfrxDiffPucchF34HalfpiBpskEnumeratedSupported
)
