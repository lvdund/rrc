package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-supportedDMRS-TypeDL ::= ENUMERATED
type PhyParametersfrxDiffSupporteddmrsTypedl struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffSupporteddmrsTypedlEnumeratedNothing = iota
	PhyParametersfrxDiffSupporteddmrsTypedlEnumeratedType1
	PhyParametersfrxDiffSupporteddmrsTypedlEnumeratedType1and2
)
