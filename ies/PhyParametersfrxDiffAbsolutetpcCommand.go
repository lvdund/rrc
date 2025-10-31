package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-absoluteTPC-Command ::= ENUMERATED
type PhyParametersfrxDiffAbsolutetpcCommand struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffAbsolutetpcCommandEnumeratedNothing = iota
	PhyParametersfrxDiffAbsolutetpcCommandEnumeratedSupported
)
