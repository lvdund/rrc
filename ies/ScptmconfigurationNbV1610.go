package ies

import "rrc/utils"

// SCPTMConfiguration-NB-v1610 ::= SEQUENCE
type ScptmconfigurationNbV1610 struct {
	ScMtchInfolistmultitbR16 ScMtchInfolistNbR14
	MultitbGapR16            *utils.ENUMERATED
	Noncriticalextension     *interface{}
}
