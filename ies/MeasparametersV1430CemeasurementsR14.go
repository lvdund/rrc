package ies

import "rrc/utils"

// MeasParameters-v1430-ceMeasurements-r14 ::= ENUMERATED
type MeasparametersV1430CemeasurementsR14 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1430CemeasurementsR14EnumeratedNothing = iota
	MeasparametersV1430CemeasurementsR14EnumeratedSupported
)
