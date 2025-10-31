package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-enhancedPowerControl-r16 ::= ENUMERATED
type PhyParametersfrxDiffEnhancedpowercontrolR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffEnhancedpowercontrolR16EnumeratedNothing = iota
	PhyParametersfrxDiffEnhancedpowercontrolR16EnumeratedSupported
)
