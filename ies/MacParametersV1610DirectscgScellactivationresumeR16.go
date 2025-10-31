package ies

import "rrc/utils"

// MAC-Parameters-v1610-directSCG-SCellActivationResume-r16 ::= ENUMERATED
type MacParametersV1610DirectscgScellactivationresumeR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1610DirectscgScellactivationresumeR16EnumeratedNothing = iota
	MacParametersV1610DirectscgScellactivationresumeR16EnumeratedSupported
)
