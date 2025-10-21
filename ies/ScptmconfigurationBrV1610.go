package ies

import "rrc/utils"

// SCPTMConfiguration-BR-v1610 ::= SEQUENCE
type ScptmconfigurationBrV1610 struct {
	ScMtchInfolistMultitbR16 ScMtchInfolistBrR14
	MultitbGapR16            *utils.ENUMERATED
	Noncriticalextension     *interface{}
}
