package ies

import "rrc/utils"

// MIMO-UE-ParametersPerTM-v1430-semiOL-r14 ::= ENUMERATED
type MimoUeParameterspertmV1430SemiolR14 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeParameterspertmV1430SemiolR14EnumeratedNothing = iota
	MimoUeParameterspertmV1430SemiolR14EnumeratedSupported
)
