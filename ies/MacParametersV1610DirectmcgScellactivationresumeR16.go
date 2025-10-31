package ies

import "rrc/utils"

// MAC-Parameters-v1610-directMCG-SCellActivationResume-r16 ::= ENUMERATED
type MacParametersV1610DirectmcgScellactivationresumeR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1610DirectmcgScellactivationresumeR16EnumeratedNothing = iota
	MacParametersV1610DirectmcgScellactivationresumeR16EnumeratedSupported
)
