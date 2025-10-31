package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-twoPUCCH-AnyOthersInSlot ::= ENUMERATED
type PhyParametersfrxDiffTwopucchAnyothersinslot struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffTwopucchAnyothersinslotEnumeratedNothing = iota
	PhyParametersfrxDiffTwopucchAnyothersinslotEnumeratedSupported
)
