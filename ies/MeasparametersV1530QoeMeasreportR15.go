package ies

import "rrc/utils"

// MeasParameters-v1530-qoe-MeasReport-r15 ::= ENUMERATED
type MeasparametersV1530QoeMeasreportR15 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1530QoeMeasreportR15EnumeratedNothing = iota
	MeasparametersV1530QoeMeasreportR15EnumeratedSupported
)
