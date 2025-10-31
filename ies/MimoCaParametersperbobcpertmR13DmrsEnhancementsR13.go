package ies

import "rrc/utils"

// MIMO-CA-ParametersPerBoBCPerTM-r13-dmrs-Enhancements-r13 ::= ENUMERATED
type MimoCaParametersperbobcpertmR13DmrsEnhancementsR13 struct {
	Value utils.ENUMERATED
}

const (
	MimoCaParametersperbobcpertmR13DmrsEnhancementsR13EnumeratedNothing = iota
	MimoCaParametersperbobcpertmR13DmrsEnhancementsR13EnumeratedDifferent
)
