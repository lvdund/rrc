package ies

import "rrc/utils"

// SL-ConfiguredGrantConfig-r16-rrc-ConfiguredSidelinkGrant-r16-sl-TimeReferenceSFN-Type1-r16 ::= ENUMERATED
type SlConfiguredgrantconfigR16RrcConfiguredsidelinkgrantR16SlTimereferencesfnType1R16 struct {
	Value utils.ENUMERATED
}

const (
	SlConfiguredgrantconfigR16RrcConfiguredsidelinkgrantR16SlTimereferencesfnType1R16EnumeratedNothing = iota
	SlConfiguredgrantconfigR16RrcConfiguredsidelinkgrantR16SlTimereferencesfnType1R16EnumeratedSfn512
)
