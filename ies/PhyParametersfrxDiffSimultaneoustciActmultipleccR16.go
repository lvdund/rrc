package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-simultaneousTCI-ActMultipleCC-r16 ::= ENUMERATED
type PhyParametersfrxDiffSimultaneoustciActmultipleccR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffSimultaneoustciActmultipleccR16EnumeratedNothing = iota
	PhyParametersfrxDiffSimultaneoustciActmultipleccR16EnumeratedSupported
)
