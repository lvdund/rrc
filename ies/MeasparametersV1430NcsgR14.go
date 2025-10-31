package ies

import "rrc/utils"

// MeasParameters-v1430-ncsg-r14 ::= ENUMERATED
type MeasparametersV1430NcsgR14 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1430NcsgR14EnumeratedNothing = iota
	MeasparametersV1430NcsgR14EnumeratedSupported
)
