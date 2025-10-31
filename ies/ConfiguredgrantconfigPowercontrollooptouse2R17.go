package ies

import "rrc/utils"

// ConfiguredGrantConfig-powerControlLoopToUse2-r17 ::= ENUMERATED
type ConfiguredgrantconfigPowercontrollooptouse2R17 struct {
	Value utils.ENUMERATED
}

const (
	ConfiguredgrantconfigPowercontrollooptouse2R17EnumeratedNothing = iota
	ConfiguredgrantconfigPowercontrollooptouse2R17EnumeratedN0
	ConfiguredgrantconfigPowercontrollooptouse2R17EnumeratedN1
)
