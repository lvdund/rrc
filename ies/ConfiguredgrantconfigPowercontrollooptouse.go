package ies

import "rrc/utils"

// ConfiguredGrantConfig-powerControlLoopToUse ::= ENUMERATED
type ConfiguredgrantconfigPowercontrollooptouse struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigPowercontrollooptouseEnumeratedNothing = iota
	ConfiguredgrantconfigPowercontrollooptouseEnumeratedN0
	ConfiguredgrantconfigPowercontrollooptouseEnumeratedN1
)
