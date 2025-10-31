package ies

import "rrc/utils"

// MeasParameters-v1310-rs-SINR-Meas-r13 ::= ENUMERATED
type MeasparametersV1310RsSinrMeasR13 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1310RsSinrMeasR13EnumeratedNothing = iota
	MeasparametersV1310RsSinrMeasR13EnumeratedSupported
)
