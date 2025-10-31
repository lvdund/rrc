package ies

import "rrc/utils"

// HighSpeedEnhParameters-r14-demodulationEnhancements-r14 ::= ENUMERATED
type HighspeedenhparametersR14DemodulationenhancementsR14 struct {
	Value utils.ENUMERATED
}

const (
	HighspeedenhparametersR14DemodulationenhancementsR14EnumeratedNothing = iota
	HighspeedenhparametersR14DemodulationenhancementsR14EnumeratedSupported
)
