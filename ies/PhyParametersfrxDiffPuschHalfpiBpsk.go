package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-pusch-HalfPi-BPSK ::= ENUMERATED
type PhyParametersfrxDiffPuschHalfpiBpsk struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffPuschHalfpiBpskEnumeratedNothing = iota
	PhyParametersfrxDiffPuschHalfpiBpskEnumeratedSupported
)
