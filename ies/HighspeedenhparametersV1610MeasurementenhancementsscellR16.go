package ies

import "rrc/utils"

// HighSpeedEnhParameters-v1610-measurementEnhancementsSCell-r16 ::= ENUMERATED
type HighspeedenhparametersV1610MeasurementenhancementsscellR16 struct {
	Value utils.ENUMERATED
}

const (
	HighspeedenhparametersV1610MeasurementenhancementsscellR16EnumeratedNothing = iota
	HighspeedenhparametersV1610MeasurementenhancementsscellR16EnumeratedSupported
)
