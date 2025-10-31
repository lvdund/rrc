package ies

import "rrc/utils"

// MeasParameters-v1530-ca-IdleModeMeasurements-r15 ::= ENUMERATED
type MeasparametersV1530CaIdlemodemeasurementsR15 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1530CaIdlemodemeasurementsR15EnumeratedNothing = iota
	MeasparametersV1530CaIdlemodemeasurementsR15EnumeratedSupported
)
