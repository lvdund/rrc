package ies

import "rrc/utils"

// MIMO-UE-ParametersPerTM-v1430-ul-dmrs-Enhancements-r14 ::= ENUMERATED
type MimoUeParameterspertmV1430UlDmrsEnhancementsR14 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeParameterspertmV1430UlDmrsEnhancementsR14EnumeratedNothing = iota
	MimoUeParameterspertmV1430UlDmrsEnhancementsR14EnumeratedSupported
)
