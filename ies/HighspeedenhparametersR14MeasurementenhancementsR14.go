package ies

import "rrc/utils"

// HighSpeedEnhParameters-r14-measurementEnhancements-r14 ::= ENUMERATED
type HighspeedenhparametersR14MeasurementenhancementsR14 struct {
	Value utils.ENUMERATED
}

const (
	HighspeedenhparametersR14MeasurementenhancementsR14EnumeratedNothing = iota
	HighspeedenhparametersR14MeasurementenhancementsR14EnumeratedSupported
)
