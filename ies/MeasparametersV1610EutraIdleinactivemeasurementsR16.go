package ies

import "rrc/utils"

// MeasParameters-v1610-eutra-IdleInactiveMeasurements-r16 ::= ENUMERATED
type MeasparametersV1610EutraIdleinactivemeasurementsR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1610EutraIdleinactivemeasurementsR16EnumeratedNothing = iota
	MeasparametersV1610EutraIdleinactivemeasurementsR16EnumeratedSupported
)
