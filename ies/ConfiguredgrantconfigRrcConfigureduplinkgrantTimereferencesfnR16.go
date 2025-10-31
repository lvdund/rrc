package ies

import "rrc/utils"

// ConfiguredGrantConfig-rrc-ConfiguredUplinkGrant-timeReferenceSFN-r16 ::= ENUMERATED
type ConfiguredgrantconfigRrcConfigureduplinkgrantTimereferencesfnR16 struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigRrcConfigureduplinkgrantTimereferencesfnR16EnumeratedNothing = iota
	ConfiguredgrantconfigRrcConfigureduplinkgrantTimereferencesfnR16EnumeratedSfn512
)
