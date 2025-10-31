package ies

import "rrc/utils"

// Phy-ParametersXDD-Diff-dynamicSFI ::= ENUMERATED
type PhyParametersxddDiffDynamicsfi struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersxddDiffDynamicsfiEnumeratedNothing = iota
	PhyParametersxddDiffDynamicsfiEnumeratedSupported
)
