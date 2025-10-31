package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-twoFL-DMRS-TwoAdditionalDMRS-UL ::= ENUMERATED
type PhyParametersfrxDiffTwoflDmrsTwoadditionaldmrsUl struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffTwoflDmrsTwoadditionaldmrsUlEnumeratedNothing = iota
	PhyParametersfrxDiffTwoflDmrsTwoadditionaldmrsUlEnumeratedSupported
)
