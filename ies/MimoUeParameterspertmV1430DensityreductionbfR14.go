package ies

import "rrc/utils"

// MIMO-UE-ParametersPerTM-v1430-densityReductionBF-r14 ::= ENUMERATED
type MimoUeParameterspertmV1430DensityreductionbfR14 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeParameterspertmV1430DensityreductionbfR14EnumeratedNothing = iota
	MimoUeParameterspertmV1430DensityreductionbfR14EnumeratedSupported
)
