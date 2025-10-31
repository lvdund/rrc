package ies

import "rrc/utils"

// HighSpeedParameters-r16-demodulationEnhancement-r16 ::= ENUMERATED
type HighspeedparametersR16DemodulationenhancementR16 struct {
	Value utils.ENUMERATED
}

const (
	HighspeedparametersR16DemodulationenhancementR16EnumeratedNothing = iota
	HighspeedparametersR16DemodulationenhancementR16EnumeratedSupported
)
