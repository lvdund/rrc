package ies

import "rrc/utils"

// HighSpeedEnhParameters-v1610-demodulationEnhancements2-r16 ::= ENUMERATED
type HighspeedenhparametersV1610Demodulationenhancements2R16 struct {
	Value utils.ENUMERATED
}

const (
	HighspeedenhparametersV1610Demodulationenhancements2R16EnumeratedNothing = iota
	HighspeedenhparametersV1610Demodulationenhancements2R16EnumeratedSupported
)
