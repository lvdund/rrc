package ies

import "rrc/utils"

// MeasParameters-v1430-nonUniformGap-r14 ::= ENUMERATED
type MeasparametersV1430NonuniformgapR14 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1430NonuniformgapR14EnumeratedNothing = iota
	MeasparametersV1430NonuniformgapR14EnumeratedSupported
)
