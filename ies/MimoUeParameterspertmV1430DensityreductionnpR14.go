package ies

import "rrc/utils"

// MIMO-UE-ParametersPerTM-v1430-densityReductionNP-r14 ::= ENUMERATED
type MimoUeParameterspertmV1430DensityreductionnpR14 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeParameterspertmV1430DensityreductionnpR14EnumeratedNothing = iota
	MimoUeParameterspertmV1430DensityreductionnpR14EnumeratedSupported
)
