package ies

import "rrc/utils"

// MeasParameters-v1430-shortMeasurementGap-r14 ::= ENUMERATED
type MeasparametersV1430ShortmeasurementgapR14 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1430ShortmeasurementgapR14EnumeratedNothing = iota
	MeasparametersV1430ShortmeasurementgapR14EnumeratedSupported
)
