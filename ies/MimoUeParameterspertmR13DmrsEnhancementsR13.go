package ies

import "rrc/utils"

// MIMO-UE-ParametersPerTM-r13-dmrs-Enhancements-r13 ::= ENUMERATED
type MimoUeParameterspertmR13DmrsEnhancementsR13 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeParameterspertmR13DmrsEnhancementsR13EnumeratedNothing = iota
	MimoUeParameterspertmR13DmrsEnhancementsR13EnumeratedSupported
)
