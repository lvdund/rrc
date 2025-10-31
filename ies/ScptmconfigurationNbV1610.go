package ies

// SCPTMConfiguration-NB-v1610 ::= SEQUENCE
type ScptmconfigurationNbV1610 struct {
	ScMtchInfolistmultitbR16 ScMtchInfolistNbR14
	MultitbGapR16            *ScptmconfigurationNbV1610MultitbGapR16
	Noncriticalextension     *ScptmconfigurationNbV1610Noncriticalextension
}
