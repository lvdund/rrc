package ies

import "rrc/utils"

// MAC-ParametersFRX-Diff-r16-directSCG-SCellActivation-r16 ::= ENUMERATED
type MacParametersfrxDiffR16DirectscgScellactivationR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersfrxDiffR16DirectscgScellactivationR16EnumeratedNothing = iota
	MacParametersfrxDiffR16DirectscgScellactivationR16EnumeratedSupported
)
