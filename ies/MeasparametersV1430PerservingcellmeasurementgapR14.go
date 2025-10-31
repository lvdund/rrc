package ies

import "rrc/utils"

// MeasParameters-v1430-perServingCellMeasurementGap-r14 ::= ENUMERATED
type MeasparametersV1430PerservingcellmeasurementgapR14 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1430PerservingcellmeasurementgapR14EnumeratedNothing = iota
	MeasparametersV1430PerservingcellmeasurementgapR14EnumeratedSupported
)
