package ies

import "rrc/utils"

// MeasParameters-v1530-qoe-MTSI-MeasReport-r15 ::= ENUMERATED
type MeasparametersV1530QoeMtsiMeasreportR15 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1530QoeMtsiMeasreportR15EnumeratedNothing = iota
	MeasparametersV1530QoeMtsiMeasreportR15EnumeratedSupported
)
