package ies

import "rrc/utils"

// MAC-ParametersFRX-Diff-r16-directSCG-SCellActivationResume-r16 ::= ENUMERATED
type MacParametersfrxDiffR16DirectscgScellactivationresumeR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersfrxDiffR16DirectscgScellactivationresumeR16EnumeratedNothing = iota
	MacParametersfrxDiffR16DirectscgScellactivationresumeR16EnumeratedSupported
)
