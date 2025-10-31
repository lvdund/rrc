package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-twoTCI-Act-servingCellInCC-List-r16 ::= ENUMERATED
type PhyParametersfrxDiffTwotciActServingcellinccListR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffTwotciActServingcellinccListR16EnumeratedNothing = iota
	PhyParametersfrxDiffTwotciActServingcellinccListR16EnumeratedSupported
)
