package ies

import "rrc/utils"

// HighSpeedParameters-r16-measurementEnhancement-r16 ::= ENUMERATED
type HighspeedparametersR16MeasurementenhancementR16 struct {
	Value utils.ENUMERATED
}

const (
	HighspeedparametersR16MeasurementenhancementR16EnumeratedNothing = iota
	HighspeedparametersR16MeasurementenhancementR16EnumeratedSupported
)
