package ies

import "rrc/utils"

// MAC-ParametersFRX-Diff-r16-directMCG-SCellActivation-r16 ::= ENUMERATED
type MacParametersfrxDiffR16DirectmcgScellactivationR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersfrxDiffR16DirectmcgScellactivationR16EnumeratedNothing = iota
	MacParametersfrxDiffR16DirectmcgScellactivationR16EnumeratedSupported
)
