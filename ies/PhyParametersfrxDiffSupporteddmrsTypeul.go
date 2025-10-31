package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-supportedDMRS-TypeUL ::= ENUMERATED
type PhyParametersfrxDiffSupporteddmrsTypeul struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffSupporteddmrsTypeulEnumeratedNothing = iota
	PhyParametersfrxDiffSupporteddmrsTypeulEnumeratedType1
	PhyParametersfrxDiffSupporteddmrsTypeulEnumeratedType1and2
)
