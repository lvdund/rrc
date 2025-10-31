package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-oneFL-DMRS-TwoAdditionalDMRS-UL ::= ENUMERATED
type PhyParametersfrxDiffOneflDmrsTwoadditionaldmrsUl struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffOneflDmrsTwoadditionaldmrsUlEnumeratedNothing = iota
	PhyParametersfrxDiffOneflDmrsTwoadditionaldmrsUlEnumeratedSupported
)
