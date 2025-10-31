package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-oneFL-DMRS-ThreeAdditionalDMRS-UL ::= ENUMERATED
type PhyParametersfrxDiffOneflDmrsThreeadditionaldmrsUl struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffOneflDmrsThreeadditionaldmrsUlEnumeratedNothing = iota
	PhyParametersfrxDiffOneflDmrsThreeadditionaldmrsUlEnumeratedSupported
)
