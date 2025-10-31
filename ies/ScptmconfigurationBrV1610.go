package ies

// SCPTMConfiguration-BR-v1610 ::= SEQUENCE
type ScptmconfigurationBrV1610 struct {
	ScMtchInfolistMultitbR16 ScMtchInfolistBrR14
	MultitbGapR16            *ScptmconfigurationBrV1610MultitbGapR16
	Noncriticalextension     *ScptmconfigurationBrV1610Noncriticalextension
}
