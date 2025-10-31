package ies

import "rrc/utils"

// MIMO-UE-ParametersPerTM-v1430-hybridCSI-r14 ::= ENUMERATED
type MimoUeParameterspertmV1430HybridcsiR14 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeParameterspertmV1430HybridcsiR14EnumeratedNothing = iota
	MimoUeParameterspertmV1430HybridcsiR14EnumeratedSupported
)
