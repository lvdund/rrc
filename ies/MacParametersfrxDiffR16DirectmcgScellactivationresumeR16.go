package ies

import "rrc/utils"

// MAC-ParametersFRX-Diff-r16-directMCG-SCellActivationResume-r16 ::= ENUMERATED
type MacParametersfrxDiffR16DirectmcgScellactivationresumeR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersfrxDiffR16DirectmcgScellactivationresumeR16EnumeratedNothing = iota
	MacParametersfrxDiffR16DirectmcgScellactivationresumeR16EnumeratedSupported
)
