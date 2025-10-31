package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-dynamicSFI ::= ENUMERATED
type PhyParametersfrxDiffDynamicsfi struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffDynamicsfiEnumeratedNothing = iota
	PhyParametersfrxDiffDynamicsfiEnumeratedSupported
)
