package ies

import "rrc/utils"

// ConfiguredGrantConfig-autonomousTx-r16 ::= ENUMERATED
type ConfiguredgrantconfigAutonomoustxR16 struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigAutonomoustxR16EnumeratedNothing = iota
	ConfiguredgrantconfigAutonomoustxR16EnumeratedEnabled
)
